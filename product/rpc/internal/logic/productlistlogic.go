package logic

import (
	"context"
	"fmt"
	"strconv"
	"time"

	"hnchain/common/xerr"
	"hnchain/product/rpc/internal/svc"
	"hnchain/product/rpc/model"
	"hnchain/product/rpc/product"

	"github.com/pkg/errors"
	"github.com/zeromicro/go-zero/core/logx"
	"github.com/zeromicro/go-zero/core/mr"
	"github.com/zeromicro/go-zero/core/threading"
)

type ProductListLogic struct {
	ctx    context.Context
	svcCtx *svc.ServiceContext
	logx.Logger
}

const (
	defaultPageSize = 10
	defaultLimit    = 300
	expireTime      = 3600 * 24 * 3
)

func NewProductListLogic(ctx context.Context, svcCtx *svc.ServiceContext) *ProductListLogic {
	return &ProductListLogic{
		ctx:    ctx,
		svcCtx: svcCtx,
		Logger: logx.WithContext(ctx),
	}
}

// 分页查询
func (l *ProductListLogic) ProductList(in *product.ProductListReq) (*product.ProductListRsp, error) {

	_, err := l.svcCtx.ProductCategoryModel.FindOne(l.ctx, in.Categoryid)
	if err == model.ErrNotFound {
		return nil, errors.Wrapf(xerr.NewErrCode(xerr.DataNoExistError), "categoryid=%d not found", in.Categoryid)
	}
	if in.Cursor == 0 {
		in.Cursor = time.Now().Unix()
	}
	if in.Ps == 0 {
		in.Ps = defaultPageSize
	}
	var (
		isCache, isEnd   bool
		lastID, lastTime int64
		firstPage        []*product.ProductItem
		products        []*model.Product
	)

	pids, _ := l.cacheProductList(l.ctx, in.Categoryid, in.Cursor, in.Ps)
	if len(pids) == int(in.Ps) {
		isCache = true
		if pids[len(pids)-1] == -1 {
			isEnd = true
		}
		products, err = l.productsByIds(l.ctx, pids)
		if err != nil {
			l.Logger.Errorf("productsByIds db err:%v", err)
			return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "productsByIds:db query err!")
		}

		for _, p := range products {
			firstPage = append(firstPage, &product.ProductItem{
				Productid:p.Id,
				Name: p.Name,
				Stock: p.Stock,
				Imageurl: p.Imageurl,
				Price: p.Price,
				Description: p.Detail,
				Categoryid: p.Categoryid,
				Status: p.Status,
				Createtime: p.Createtime.Unix(),
			})
		}
	} else {
		
		products, err = l.svcCtx.ProductModel.CategoryProducts(l.ctx, in.Cursor, in.Categoryid, defaultLimit)
		if err != nil {
			return nil, errors.Wrap(xerr.NewErrCode(xerr.DbError), "CategoryProducts: db query err!")
		}
		var firstPageProducts []*model.Product
		if len(products) > int(in.Ps) {
			firstPageProducts = products[0:int(in.Ps)]
		} else {
			firstPageProducts = products
			isEnd = true
		}

		for _, p := range firstPageProducts {
			firstPage = append(firstPage, &product.ProductItem{
				Productid:p.Id,
				Name: p.Name,
				Stock: p.Stock,
				Imageurl: p.Imageurl,
				Price: p.Price,
				Description: p.Detail,
				Categoryid: p.Categoryid,
				Status: p.Status,
				Createtime: p.Createtime.Unix(),
			})
		}
	}

	if len(firstPage) > 0 {
		pageLast := firstPage[len(firstPage)-1]
		lastID = pageLast.Productid
		lastTime = pageLast.Createtime
		if lastTime < 0 {
			lastTime = 0
		}

		//去重复
		for k, p := range firstPage {
			if p.Createtime == in.Cursor && p.Productid == in.Productid {
				firstPage = firstPage[k:]
				break
			}
		}
	}

	ret := &product.ProductListRsp{
		IsEnd:     isEnd,
		Timestamp: lastTime,
		Productid: lastID,
		Products:  firstPage,
	}

	if !isCache {
		threading.GoSafe(func() {
			if len(products) < defaultLimit && len(products) > 0 {
				endTime, _ := time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
				products = append(products, &model.Product{Id: -1,Createtime: endTime})
			}
			_ = l.addCacheProductList(context.Background(), products)
		})
	}

	return ret, nil
}

/*
* categoryId 分类Id
* cursor 查询时间
* ps 页大小
 */

func (l *ProductListLogic) cacheProductList(ctx context.Context, categoryId int64, cursor int64, ps int32) ([]int64, error) {

	//倒序
	pairs, err := l.svcCtx.BizRedis.ZrevrangebyscoreWithScoresAndLimitCtx(ctx,categoryKey(categoryId),cursor,0, 0,int(ps))
	if err != nil {
		return nil, err
	}
	var ids []int64
	for _, pair := range pairs {
		id, _ := strconv.ParseInt(pair.Key, 10, 64)
		ids = append(ids, id)
	}
	return ids, nil
}
func (l *ProductListLogic) addCacheProductList(ctx context.Context, products []*model.Product) error {
	if len(products) == 0 {
		return nil
	}
	for _, p := range products {
		score := p.Createtime.Unix()
		if score < 0 {
			score = 0
		}
		_, err := l.svcCtx.BizRedis.ZaddCtx(ctx, categoryKey(p.Categoryid), score, strconv.FormatInt(p.Id, 10))
		if err != nil {
			return err
		}
	}
	return l.svcCtx.BizRedis.ExpireCtx(ctx, categoryKey(products[0].Categoryid), expireTime)
}

func categoryKey(categoryId int64) string {
	return fmt.Sprintf("product:categoryKey:%d", categoryId)
}

func (l *ProductListLogic) productsByIds(ctx context.Context, pids []int64) ([]*model.Product, error) {
	products, err := mr.MapReduce(func(source chan<- any) {
		for _, p := range pids {
			source <- p
		}
	},
	func(it any, writer mr.Writer[any], cancal func(error)) {
		id := it.(int64)
        if id == -1 {
			endTime, _ := time.Parse("2006-01-02 15:04:05", "0000-00-00 00:00:00")
			writer.Write(&model.Product{
				Id: id,
				Createtime: endTime,
			})
			return
		}

		p, err := l.svcCtx.ProductModel.FindOne(ctx, id)
		if err != nil {
			cancal(err)
			return
		}
		writer.Write(p)
	},
	func(pipe <-chan any, writer mr.Writer[any], cancel func(error)) {
		mp :=make(map[int64]*model.Product,0)
		for it := range pipe {
			v :=it.(*model.Product)
			mp[v.Id]=v
		}
		writer.Write(mp)
	})
	
	if err != nil {
		return nil, err
	}
	var rt  []*model.Product
    mpProduts := products.(map[int64]*model.Product)  
	for _,it := range pids {
		if v,ok := mpProduts[it];ok {
			rt =append(rt, v)
		}
	}

	return rt, nil
}
