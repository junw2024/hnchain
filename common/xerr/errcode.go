package xerr
//成功返回
const OK uint32 = 200

/**(前3位代表业务,后三位代表具体功能)**/

/**全局错误码*/
//服务器开小差
const ServerCommonError uint32 = 100001

//请求参数错误
const ReuqestParamError uint32 = 100002

//token过期
const TokenExpireError uint32 = 100003

//生成token失败
const TokenGenerateError uint32 = 100004

//数据库繁忙,请稍后再试
const DbError uint32 = 100005

//更新数据影响行数为0
const DbUpdateAffectedZeroError uint32 = 100006

//数据不存在
const DataNoExistError uint32 = 100007

//Redis异常
const RedisError uint32 = 100008

//事务错误
const  TransactionError uint32 = 100009

//用户服务

//订单服务
const OrderCreateError uint32 = 300001  //创建订单错误
const OrderRevertError uint32 = 300002  //创建订单回撤错误

//商品服务
const ProductExistError   uint32 = 200001  //商品不存在
const StockInsufficient    uint32 = 200002  //库存不足

//支付服务