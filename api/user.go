package api

import (
	"encoding/json"
	"zeh/MyGoFramework/base/iface"
)

type UserInfo struct {
}

func (u *UserInfo) PreHandle(request iface.IRequest) {
	//TODO implement me
	//panic("implement me")
}

func (u *UserInfo) Handle(request iface.IRequest) {
	//TODO implement me
	data, err := json.Marshal(request.GetConnection().GetAllProperty())
	if err != nil {
		request.GetResponse().SendBuffer([]byte("用户信息缺失"))
		return
	}

	request.GetResponse().SendBuffer(data)
}

func (u *UserInfo) PostHandle(request iface.IRequest) {
	//TODO implement me
	//panic("implement me")
}
