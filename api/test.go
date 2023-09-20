package api

import (
	"fmt"
	"zeh/MyGoFramework/base/iface"
)

type Test struct {
}

func (c *Test) PreHandle(request iface.IRequest) {
	//TODO implement me
	//fmt.Println("业务处理前")
}

func (c *Test) Handle(request iface.IRequest) {
	//TODO implement me
	//fmt.Println("业务处理中")

	res := "Hello World!"
	fmt.Printf("--->客户端请求数据:%s\n", string(request.GetData()))
	//fmt.Printf("--->服务端响应数据:%s\n", res)
	request.GetResponse().Send(0, []byte(res))
}

func (c *Test) PostHandle(request iface.IRequest) {
	//TODO implement me
	//fmt.Println("业务处理后")
}
