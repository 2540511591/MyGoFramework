package iface

import (
	"context"
)

type IServer interface {
	Start()
	Stop()

	GetPack() func(IMessage, IConnection) ([]byte, error)
	GetUnPack() func(IConnection) (IMessage, error)
	GetCtx() context.Context
	GetMsgHandle() IHandle

	AddRouter(uint32, IRouter)
}
