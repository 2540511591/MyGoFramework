package base

import "zeh/MyGoFramework/base/iface"

type RouterHandle struct {
	router    iface.IRouter
	pipelines func(iface.IRequest) iface.IResponse
	req       iface.IRequest
}

func (h *RouterHandle) SetRequest(request iface.IRequest) {
	//TODO implement me
	h.req = request
}

func (h *RouterHandle) GetRouter() iface.IRouter {
	//TODO implement me
	return h.router
}

func (h *RouterHandle) Handle() iface.IResponse {
	//TODO implement me
	h.req.BindRouter(h.router)
	return h.pipelines(h.req)
}

func NewRouterHandle(router iface.IRouter, f func(iface.IRequest) iface.IResponse) iface.IRouterHandle {
	h := &RouterHandle{
		router:    router,
		pipelines: f,
	}

	return h
}
