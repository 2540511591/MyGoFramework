package api

import (
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
	request.GetResponse().Send(0, []byte(res))
}

func (c *Test) PostHandle(request iface.IRequest) {
	//TODO implement me
	//fmt.Println("业务处理后")
}
