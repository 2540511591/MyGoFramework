package base

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
	"zeh/MyGoFramework/utils"
)

type Handle struct {
	//路由集合
	api map[uint32]iface.IRouter
	//工作协程池
	workers []chan iface.IRequest
	server  iface.IServer
	//管道
	pipeline func(iface.IRequest) iface.IResponse
}

// 将请求分配给一个工作协程进行处理
func (h *Handle) Execute(request iface.IRequest) {
	//TODO implement me
	len := uint64(len(h.workers))
	id := uint32(request.GetConnection().GetId() % len)

	h.workers[id] <- request
}

func (h *Handle) AddHandle(u uint32, router iface.IRouter) {
	//TODO implement me
	h.api[u] = router
}

// 开启协程池
func (h *Handle) StartWorkerPool() {
	//TODO implement me
	fmt.Printf("####协程池启动中，协程数量：%d，单个工作协程任务数量：%d\n", conf.ServerConfig.WorkerNumber, conf.ServerConfig.WorkerQueueLen)
	var i uint32 = 0
	for ; i < conf.ServerConfig.WorkerNumber; i++ {
		h.workers[i] = make(chan iface.IRequest, conf.ServerConfig.WorkerQueueLen)

		go h.startOneWorker(i, h.workers[i])
	}
}

// 阻塞，单个协程的业务处理逻辑
func (h *Handle) startOneWorker(workerId uint32, queue chan iface.IRequest) {
	for {
		select {
		case <-h.server.GetCtx().Done():
			return
		case req := <-queue:
			router := h.api[req.GetMessageId()]
			if router == nil {
				fmt.Printf("!!!!路由为空,routerId:%d\n", req.GetMessageId())
				continue
			}

			//通过反射实例化新的路由并绑定
			req.BindRouter(utils.NewObject(router))

			//处理
			utils.Try(func() {
				_ = h.pipeline(req).Output()
			}, func(err interface{}) {
				fmt.Printf("!!!!请求处理异常,err:%s\n", err.(error))
				return
			})
		}
	}
}

// 最后一个管道逻辑
func final(r iface.IRequest) iface.IResponse {
	r.Call()
	return r.GetResponse()
}

func NewHandle(server iface.IServer) iface.IHandle {
	h := &Handle{
		api:     make(map[uint32]iface.IRouter),
		workers: make([]chan iface.IRequest, conf.ServerConfig.WorkerNumber),
		server:  server,
	}

	//初始化管道
	pipe := &utils.PipeLine[iface.IRequest, iface.IResponse]{}
	//设置管道列表
	pipe.SetPipes(conf.CommonPipeline)
	//设置最后的管道
	pipe.SetFinal(final)
	//创建
	h.pipeline, _ = pipe.Create()

	return h
}
