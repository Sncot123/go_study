package tcpiface

type IServer interface {
	//启动服务器
	Start()
	//终止服务器
	Stop()
	//运行服务器
	Server()
	//路由功能：给当前的服务注册一个路由方法，共客户端的链接处理使用
	AddRouter(router IRouter)
}
