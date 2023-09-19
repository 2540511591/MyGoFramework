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
	fmt.Println(1)
	res := next(request)
	fmt.Println(1)

	return res
}

func Test2(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println(2)
	res := next(request)
	fmt.Println(2)

	return res
}

func Test3(request iface.IRequest, next func(iface.IRequest) iface.IResponse) iface.IResponse {
	fmt.Println(3)
	res := next(request)
	fmt.Println(3)

	return res
}
