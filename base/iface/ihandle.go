package iface

type IHandle interface {
	//执行一个请求
	Execute(IRequest)
	//开启工作协程池
	StartWorkerPool()
}
