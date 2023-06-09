package tcpiface

/*
	·根据消息调度路由处理和添加接口
*/
type IMsgHandle interface {
	//调度/执行对应的router处理方法
	DoMsgHandle(request IRequest)

	//为消息添加具体的处理逻辑
	AddRouter(msgID uint32, router IRouter)
}
