package iface

type IRouter interface {
	//业务处理前
	PreHandle(IRequest)
	//业务处理中
	Handle(IRequest)
	//业务处理后
	PostHandle(IRequest)
}
