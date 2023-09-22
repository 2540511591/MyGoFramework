package base

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
	"zeh/MyGoFramework/utils"
)

type Handle struct {
	//工作协程池
	workers      []chan iface.IRequest
	server       iface.IServer
	routerManage iface.IRouterManager
}

// 将请求分配给一个工作协程进行处理
func (h *Handle) Execute(request iface.IRequest) {
	//TODO implement me
	len := uint64(len(h.workers))
	id := uint32(request.GetConnection().GetId() % len)

	h.workers[id] <- request
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
			routerH, err := h.routerManage.GetRouter(req.GetMessageId())
			if err != nil {
				fmt.Printf("!!!!路由为空,routerId:%d\n", req.GetMessageId())
				continue
			}

			routerH.SetRequest(req)

			//处理
			utils.Try(func() {
				_ = routerH.Handle().Output()
			}, func(err interface{}) {
				fmt.Printf("!!!!请求处理异常,err:%s\n", err.(error))
				return
			})
		}
	}
}

func NewHandle(server iface.IServer) iface.IHandle {
	h := &Handle{
		workers:      make([]chan iface.IRequest, conf.ServerConfig.WorkerNumber),
		server:       server,
		routerManage: server.GetRouterManage(),
	}

	return h
}
