package conf

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
)

var (
	//公共管道
	CommonPipeline = []func(iface.IRequest, func(iface.IRequest) iface.IResponse) iface.IResponse{
		Test1,
		Test2,
		Test3,
	}
)

func Test1(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	if request.GetMessageId() == 1 {
		fmt.Println("请求路由id为1，拦截")
		request.GetResponse().SendBuffer([]byte("非法请求，请重试！"))
		return request.GetResponse()
	}

	res := next(request)

	return res
}

func Test2(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	res := next(request)

	return res
}

func Test3(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	res := next(request)

	return res
}
