package base

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
	"zeh/MyGoFramework/conf"
	"zeh/MyGoFramework/utils"
)

type RouterGroup struct {
	name      string
	manage    iface.IRouterManager
	pipelines []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse
}

func (g *RouterGroup) GetName() string {
	//TODO implement me
	return g.name
}

func (g *RouterGroup) AddRouter(u uint32, router iface.IRouter) {
	//TODO implement me
	g.manage.AddRouter(u, router)
	g.manage.SetRouterGroup(u, g.name)
}

func (g *RouterGroup) AddRouterPl(u uint32, router iface.IRouter, f func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) {
	//TODO implement me
	g.AddRouter(u, router)
	g.manage.SetRouterPipeLine(u, f)
}

func (g *RouterGroup) Group(s string) iface.IRouterGroup {
	//TODO implement me
	name := fmt.Sprintf("%s.%s", g.name, s)
	var pipes []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse
	pipes = append(pipes, g.pipelines...)
	if v, ok := conf.GroupPipeline[name]; ok {
		pipes = append(pipes, v...)
	}

	return NewRouterGroup(g.manage, name, pipes)
}

func (g *RouterGroup) GroupPl(s string, f func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) iface.IRouterGroup {
	//TODO implement me
	name := fmt.Sprintf("%s.%s", g.name, s)
	var pipes []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse
	pipes = append(pipes, g.pipelines...)
	if v, ok := conf.GroupPipeline[name]; ok {
		pipes = append(pipes, v...)
	}
	pipes = append(pipes, f)

	return NewRouterGroup(g.manage, name, pipes)
}

func NewRouterGroup(manage iface.IRouterManager, name string, pipelines []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse) iface.IRouterGroup {
	g := &RouterGroup{
		name:      name,
		manage:    manage,
		pipelines: pipelines,
	}

	pipe := &utils.PipeLine[iface.IRequest, iface.IResponse]{}
	pipe.SetPipes(pipelines)
	pipe.SetFinal(manage.GetPipeFinal())
	pipes, _ := pipe.Create()
	manage.SetGroupPipeLines(name, pipes)

	return g
}
