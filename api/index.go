package api

import "zeh/MyGoFramework/base/iface"

type Index struct {
}

func (c *Index) PreHandle(request iface.IRequest) {
	//TODO implement me
	request.GetResponse().Send(0, []byte("你好"))
}

func (c *Index) Handle(request iface.IRequest) {
	//TODO implement me
	request.GetResponse().Send(0, []byte("你好"))
}

func (c *Index) PostHandle(request iface.IRequest) {
	//TODO implement me
	request.GetResponse().Send(0, []byte("你好"))
}
