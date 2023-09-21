package base

import (
	"errors"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
	"zeh/MyGoFramework/utils"
)

type RouterManage struct {
	//所有的路由
	routers map[uint32]iface.IRouter
	//路由绑定的组
	bindingGroups map[uint32]string
	//组的管道
	groupPipelines map[string]func(iface.IRequest) iface.IResponse

	//路由单独对应的管道
	routerPipeline map[uint32]func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse

	// 公共管道
	commonPipelines func(iface.IRequest) iface.IResponse
}

func (m *RouterManage) SetRouterPipeLine(u uint32, f func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) {
	//TODO implement me
	m.routerPipeline[u] = f
}

func (m *RouterManage) SetRouterGroup(u uint32, s string) {
	//TODO implement me
	m.bindingGroups[u] = s
}

func (m *RouterManage) GetPipeFinal() func(iface.IRequest) iface.IResponse {
	//TODO implement me
	return func(request iface.IRequest) iface.IResponse {
		f := func(iRequest iface.IRequest) iface.IResponse {
			iRequest.Call()
			return iRequest.GetResponse()
		}

		routerPipe := m.routerPipeline[request.GetMessageId()]
		if routerPipe != nil {
			return routerPipe(request, f)
		}

		return f(request)
	}
}

func (m *RouterManage) GetName() string {
	//TODO implement me
	return ""
}

func (m *RouterManage) AddRouter(u uint32, router iface.IRouter) {
	//TODO implement me
	m.routers[u] = router
}

func (m *RouterManage) AddRouterPl(u uint32, router iface.IRouter, f func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) {
	//TODO implement me
	m.AddRouter(u, router)
	m.routerPipeline[u] = f
}

func (m *RouterManage) Group(s string) iface.IRouterGroup {
	//TODO implement me
	var pipes []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse
	pipes = append(pipes, conf.CommonPipeline...)
	if v, ok := conf.GroupPipeline[s]; ok {
		pipes = append(pipes, v...)
	}

	return NewRouterGroup(m, s, pipes)
}

func (m *RouterManage) GroupPl(s string, f func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) iface.IRouterGroup {
	//TODO implement me
	var pipes []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse
	pipes = append(pipes, conf.CommonPipeline...)
	if v, ok := conf.GroupPipeline[s]; ok {
		pipes = append(pipes, v...)
	}
	pipes = append(pipes, f)

	return NewRouterGroup(m, s, pipes)
}

func (m *RouterManage) GetRouter(u uint32) (iface.IRouterHandle, error) {
	//TODO implement me
	//查找路由
	router := m.routers[u]
	if router == nil {
		return nil, errors.New("路由不存在")
	}
	router = utils.NewObject(router)

	pipes := m.commonPipelines
	// 查找组
	group := m.bindingGroups[u]
	if group != "" {
		groupPipes := m.groupPipelines[group]
		if groupPipes != nil {
			pipes = groupPipes
		}
	}

	return NewRouterHandle(router, pipes), nil
}

func (m *RouterManage) SetGroupPipeLines(s string, f func(iface.IRequest) iface.IResponse) {
	//TODO implement me
	m.groupPipelines[s] = f
}

func NewRouterManage() iface.IRouterManager {
	m := &RouterManage{
		routers:         make(map[uint32]iface.IRouter),
		bindingGroups:   make(map[uint32]string),
		groupPipelines:  make(map[string]func(iface.IRequest) iface.IResponse),
		routerPipeline:  make(map[uint32]func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse),
		commonPipelines: nil,
	}

	pipe := &utils.PipeLine[iface.IRequest, iface.IResponse]{}
	pipe.SetPipes(conf.CommonPipeline)
	pipe.SetFinal(m.GetPipeFinal())
	pipes, _ := pipe.Create()
	m.commonPipelines = pipes

	return m
}
