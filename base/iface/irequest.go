package iface

type IRequest interface {
	//获取消息对象
	GetMessage() IMessage
	//获取消息id
	GetMessageId() uint32
	//获取消息数据
	GetData() []byte
	//获取客户端连接对象
	GetConnection() IConnection
	//获取响应
	GetResponse() IResponse

	//绑定路由
	BindRouter(IRouter)
	//执行
	Call()
}
