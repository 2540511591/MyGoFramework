package iface

type IRequest interface {
	GetMessage() IMessage
	GetMessageId() uint32
	GetData() []byte
	GetConnection() IConnection
	GetResponse() IResponse

	BindRouter(IRouter)
	Call()
}
