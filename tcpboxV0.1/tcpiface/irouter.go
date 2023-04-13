package tcpiface

/*
	路由抽象接口
	路由里的数据都是IRequest
*/

type IRouter interface {
	//处理conn业务之前的钩子方法Hook
	PreHandle(request IRequest)

	//处理conn的主方法hook
	Handle(request IRequest)

	//处理conn业务之后的钩子方法Hook
	PostHandle(request IRequest)
}
