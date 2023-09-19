package iface

type IResponse interface {
	GetMessage() IMessage
	GetMessageId() uint32
	GetData() []byte
	GetConnection() IConnection
	GetRequest() IRequest

	//设置数据，不会立即输出，需要经过管道
	Send(uint32, []byte)
	SendBuffer([]byte)
	SendMsg(IMessage)

	//立刻输出数据，不会经过管道
	Output() error
	//清空数据
	Refresh()
	//数据是否已输出
	HasOutput() bool
}
