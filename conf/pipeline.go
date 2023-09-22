package conf

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
)

var (
	//公共管道
	CommonPipelines = []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse{
		CommonPipeline,
	}

	//分组管道，对应路由分组
	GroupPipelines = map[string][]func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse{
		"api": {
			ApiPipeline,
		},
		"api.auth": {
			ApiAuthPipeline,
		},
	}
)

func CommonPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println("--->CommonPipeline")

	res := next(request)

	fmt.Println("<---CommonPipeline")

	return res
}

func ApiPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println("--->ApiPipeline")

	res := next(request)

	fmt.Println("<---ApiPipeline")

	return res
}

func ApiAuthPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println("--->ApiAuthPipeline")

	res := next(request)

	fmt.Println("<---ApiAuthPipeline")

	return res
}
