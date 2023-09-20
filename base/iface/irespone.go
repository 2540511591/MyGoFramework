package iface

type IResponse interface {
	GetConnection() IConnection
	GetRequest() IRequest
	//获取需要发送的消息数量
	GetMessageNumber() uint32
	//根据id获取消息
	GetMessage(uint32) IMessage

	//立即发送消息
	Send(uint32, []byte)
	SendBuffer([]byte)
	SendMsg(IMessage)

	//设置待发送消息,出管道后才发送
	WaitSend(uint32, []byte) uint32
	WaitBuffer([]byte) uint32
	WaitSendMsg(IMessage) uint32

	//立刻向客户端发送所有数据
	Output() error
	//清空数据
	Refresh()
}
