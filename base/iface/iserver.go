package iface

import (
	"context"
)

type IServer interface {
	Start()
	Stop()

	//获取封包函数
	GetPack() func(IMessage, IConnection) ([]byte, error)
	//获取解包函数
	GetUnPack() func(IConnection) (IMessage, error)
	//获取服务器的context
	GetCtx() context.Context
	//获取消息处理对象
	GetMsgHandle() IHandle

	//添加路由
	AddRouter(uint32, IRouter)
}
