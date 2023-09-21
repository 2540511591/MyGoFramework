package conf

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
)

var (
	//公共管道
	CommonPipeline = []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse{
		TestPipeline,
	}

	//分组管道，对应路由分组
	GroupPipeline = map[string][]func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse{
		"api": {
			TestApi,
		},
		"api.user": {
			TestUser,
		},
	}
)

func TestPipeline(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Printf("--->客户端请求，客户端ID：%d，请求路由ID：%d，请求消息：%s\n", request.GetConnection().GetId(), request.GetMessageId(), string(request.GetData()))

	res := next(request)

	fmt.Printf("<---服务端响应，客户端ID：%d，响应消息数量：%d\n", res.GetConnection().GetId(), res.GetMessageNumber())

	return res
}

func TestApi(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println("--->TestApi请求")
	res := next(request)
	fmt.Println("--->TestApi响应")

	return res
}

func TestUser(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println("--->TestUser请求")
	res := next(request)
	fmt.Println("--->TestUser响应")

	return res
}
