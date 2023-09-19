package base

import (
	"fmt"
	"reflect"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
	"zeh/MyGoFramework/utils"
)

type Handle struct {
	api      map[uint32]iface.IRouter
	workers  []chan iface.IRequest
	server   iface.IServer
	pipeline func(iface.IRequest) iface.IResponse
}

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

func (h *Handle) StartWorkerPool() {
	//TODO implement me
	fmt.Printf("--->协程池启动中，协程数量：%d，单个工作协程任务数量：%d\n", conf.ServerConfig.WorkerNumber, conf.ServerConfig.WorkerQueueLen)
	var i uint32 = 0
	for ; i < conf.ServerConfig.WorkerNumber; i++ {
		h.workers[i] = make(chan iface.IRequest, conf.ServerConfig.WorkerQueueLen)

		go h.startOneWorker(i, h.workers[i])
	}
}

func (h *Handle) startOneWorker(workerId uint32, queue chan iface.IRequest) {
	for {
		select {
		case <-h.server.GetCtx().Done():
			return
		case req := <-queue:
			router := h.api[req.GetMessageId()]
			if router == nil {
				fmt.Printf("--->路由为空,routerId:%d\n", req.GetMessageId())
				continue
			}

			//通过反射实例化新的路由并绑定
			req.BindRouter(newRouter(router))

			//处理
			utils.Try(func() {
				_ = h.pipeline(req).Output()
			}, func(err interface{}) {
				fmt.Printf("--->请求处理异常,err:%s\n", err.(error))
				return
			})
		}
	}
}

// 管道组装
func packPipeline(
	current func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse,
	prev func(request iface.IRequest) iface.IResponse,
) func(iface.IRequest) iface.IResponse {
	return func(request iface.IRequest) iface.IResponse {
		return current(request, prev)
	}
}

func final(r iface.IRequest) iface.IResponse {
	r.Call()
	return r.GetResponse()
}

func newRouter[T interface{}](s T) T {
	t := reflect.TypeOf(s).Elem()
	return reflect.New(t).Interface().(T)
}

func NewHandle(server iface.IServer) iface.IHandle {
	h := &Handle{
		api:     make(map[uint32]iface.IRouter),
		workers: make([]chan iface.IRequest, conf.ServerConfig.WorkerNumber),
		server:  server,
	}

	//组装管道
	var pipelines = final
	for _, pipeline := range conf.CommonPipeline {
		pipelines = packPipeline(pipeline, pipelines)
	}
	h.pipeline = pipelines

	return h
}
