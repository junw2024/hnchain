package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"hnchain/common/xerr"
	"hnchain/order/rpc/internal/svc"
	"hnchain/order/rpc/model"
	"hnchain/order/rpc/order"

	"github.com/jinzhu/copier"
	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

type OrdersLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

func NewOrdersLogic(ctx context.Context, svcCtx *svc.ServiceContext) *OrdersLogic {
	return &OrdersLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

const (
	defaultPageSize = 10
	defaultLimit    = 300
	expireTime      = 60*10
)

// 个人订单分页
func (l *OrdersLogic) Orders(in *order.OrdersReq) (*order.OrdersRsp, error) {

	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}
	if in.Ps == 0 {
		in.Ps = defaultPageSize
	}

	//没有错误处理
	pids,err := l.cacheProductList(l.ctx,in.Userid,in.Cursor,in.Ps)
	if err != nil {
		l.Logger.Info("cacheProductList err:%v",err)
	}

	var (
		isCache, isEnd   bool
		lastID,  lastTime int64
		firstPage        []*order.Orderitem
		orderitems      []*model.Orderitem
	)

	//足够一页大小
	if len(pids) == int(in.Ps) {
		isCache =true
		if pids[len(pids)-1] == -1 {
			isEnd =true
		}

		orderitems,err = l.orderitemsByIds(l.ctx,pids)
		if err != nil {
			l.Logger.Errorf("orderitemsByIds db err:%v", err)
			return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "orderitemsByIds:db query err!")
		}

		for _,p := range orderitems {
		
			var dst  order.Orderitem
			_=copier.Copy(&dst,p)
			
			dst.CreateTime = p.Createtime.Unix()
			dst.UpdateTime = p.Updatetime.Unix()
			
			firstPage = append(firstPage, &dst)
		}
	}else {
		orderitems,err := l.svcCtx.OrderitemModel.UserOrderitems(l.ctx,in.Userid,time.Unix(in.Cursor,0),defaultLimit)
		if err != nil {
			l.Logger.Errorf("UserOrderitems err:%v",err)
			return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "UserOrderitems: db query err!")
		}

		var firstPageOderitmes []*model.Orderitem
		if(len(orderitems)> int(in.Ps)) {
			firstPageOderitmes = orderitems[0 : int(in.Ps)]
		}else {
			firstPageOderitmes = orderitems
			isEnd =true
		}

		for _,p := range firstPageOderitmes {
			var dst  order.Orderitem
			_=copier.Copy(&dst,p)
			firstPage = append(firstPage, &dst)
		}	
	}

	if len(firstPage) >0 {
		pageLast := firstPage[len(firstPage)-1]
		lastID = pageLast.Id
		lastTime = pageLast.CreateTime

		if lastTime < 0 {
			lastTime = 0
		}

		//去重复
		for k, p := range firstPage {
			if p.CreateTime == in.Cursor && p.Id == in.Orderitemid {
				firstPage = firstPage[k:]
				break
			}
		}
	}
	ret := &order.OrdersRsp{Orderitems:firstPage,
		IsEnd: isEnd, 
		LastTime: lastTime,
		LastId: lastID}

	if !isCache	{
		threading.GoSafe(func() {
			if len(orderitems) < defaultLimit && len(orderitems)>0 {
				endTime, _ := time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
				orderitems = append(orderitems, &model.Orderitem{Id: -1,Createtime: endTime})
			}
			_ = l.addCacheOrderitimetList(context.Background(),orderitems)
		})
	}
	
	return ret,nil	

}

/*
* userid 用户ID
* cursor 查询时间
* ps 页大小
 */
func (l *OrdersLogic) cacheProductList(ctx context.Context, userid int64, cursor int64, ps int32) ([]int64, error) {
	//倒序
	pairs, err := l.svcCtx.BizRedis.ZrangebyscoreWithScoresAndLimit(orderKey(userid), cursor, 0, 0,int(ps))
	if err != nil {
		return nil,err
	}
	var ids []int64
	for _,pair := range pairs {
		id,_ := strconv.ParseInt(pair.Key,10,64)
		ids=append(ids, id)
	}
	return ids, nil
}

func(l *OrdersLogic) orderitemsByIds(ctx context.Context, pids []int64) ([]*model.Orderitem,error) {
	orderitems,err := mr.MapReduce(func(source chan<- int64) {
		for _,p := range pids {
			if p != -1 {
				source <- p
			}
		}
	},
	func(it int64, writer mr.Writer[*model.Orderitem], cancal func(error)) {
		p,err := l.svcCtx.OrderitemModel.FindOne(l.ctx,it)
		if err != nil {
			cancal(err)
			return
		}
		writer.Write(p)
	},
	func(pipe <-chan *model.Orderitem, writer mr.Writer[[]*model.Orderitem], cancel func(error)) {
		var ps []*model.Orderitem
		for it := range pipe {
			ps = append(ps, it)
		}
		writer.Write(ps)
	})

	if err != nil {
		return nil,err
	}

	return orderitems,nil
}

func (l *OrdersLogic)addCacheOrderitimetList(ctx context.Context,ordertimes []*model.Orderitem) error{
	if len(ordertimes) == 0 {
		return nil
	}
	for _,p := range ordertimes {
		score := p.Createtime.Unix()
		if score < 0 {
			score=0
		}
		_,err := l.svcCtx.BizRedis.ZaddCtx(ctx,orderKey(p.Userid),score,strconv.FormatInt(p.Id,10))
		if err != nil {
			return err
		}
	}
	return l.svcCtx.BizRedis.ExpireCtx(ctx,orderKey(ordertimes[0].Userid),expireTime)
}

func orderKey(userid int64) string {
	return fmt.Sprintf("order:ordersKey:%d", userid) 
}
