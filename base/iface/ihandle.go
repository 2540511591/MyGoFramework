package iface

type IHandle interface {
	Execute(IRequest)
	StartWorkerPool()
	AddHandle(uint32, IRouter)
}
