package iface

type IRouter interface {
	//业务处理前
	PreHandle(IRequest)
	//业务处理中
	Handle(IRequest)
	//业务处理后
	PostHandle(IRequest)
}

type IRouterHandle interface {
	GetRouter() IRouter
	SetRequest(IRequest)
	Handle() IResponse
}

type IRouterGroup interface {
	GetName() string
	AddRouter(uint32, IRouter)
	AddRouterPl(uint32, IRouter, func(IRequest, func(IRequest) IResponse) IResponse)
	Group(string) IRouterGroup
	GroupPl(string, func(IRequest, func(IRequest) IResponse) IResponse) IRouterGroup
}

type IRouterManager interface {
	IRouterGroup
	GetRouter(uint32) (IRouterHandle, error)
	GetPipeFinal() func(IRequest) IResponse
	SetRouterGroup(uint32, string)
	SetGroupPipeLines(string, func(IRequest) IResponse)
	SetRouterPipeLine(uint32, func(IRequest, func(IRequest) IResponse) IResponse)
}
