package iface

type IResponse interface {
	GetConnection() IConnection
	GetRequest() IRequest
	//获取需要发送的消息数量
	GetMessageNumber() uint32
	//根据id获取消息
	GetMessage(uint32) IMessage

	//设置消息，不会立即输出，需要经过管道
	Send(uint32, []byte)
	SendBuffer([]byte)
	SendMsg(IMessage)

	//立刻向客户端发送所有数据
	Output() error
	//清空数据
	Refresh()
}
