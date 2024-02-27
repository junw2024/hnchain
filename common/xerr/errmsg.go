package xerr

var message map[uint32]string

func init() {
	message = make(map[uint32]string)
	message[OK] = "SUCCESS"
	message[ServerCommonError] = "服务器开小差啦,稍后再来试一试"
	message[ReuqestParamError] = "参数错误"
	message[TokenExpireError] = "token失效，请重新登陆"
	message[TokenGenerateError] = "生成token失败"
	message[DbError] = "数据库繁忙,请稍后再试"
	message[DbUpdateAffectedZeroError] = "更新数据影响行数为0"
	message[DataNoExistError] = "数据不存在"
	message[RedisError] = "redis操作异常"
	message[TransactionError] = "本地事务异常"
	//商品服务
	message[ProductExistError] = "商品不存在"
	
	//订单服务
	message[OrderCreateError] = "创建订单错误"
	message[OrderRevertError] = "创建订单回撤错误"
}

func MapErrMsg(errCode uint32) string {
	if msg, ok := message[errCode]; ok {
		return msg
	} else {
		return "服务器开小差啦,稍后再来试一试"
	}
}

func IsCodeErr(errCode uint32) bool {
	if _,ok := message[errCode];ok {
		return true
	} else {
		return false
	}
}
