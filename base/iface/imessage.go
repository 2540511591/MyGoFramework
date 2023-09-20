package iface

type IMessage interface {
	//获取消息id
	GetId() uint32
	//获取消息长度
	GetLen() uint32
	//获取数据
	GetData() []byte

	SetId(uint32)
	SetLen(uint32)
	SetData([]byte)
}

type BaseMessage struct {
	Id   uint32
	Len  uint32
	Data []byte
}

func (m *BaseMessage) GetId() uint32 {
	//TODO implement me
	return m.Id
}

func (m *BaseMessage) GetLen() uint32 {
	//TODO implement me
	if m.Len != 0 {
		return m.Len
	}
	return uint32(len(m.Data))
}

func (m *BaseMessage) GetData() []byte {
	//TODO implement me
	return m.Data
}

func (m *BaseMessage) SetId(id uint32) {
	//TODO implement me
	m.Id = id
}

func (m *BaseMessage) SetLen(len uint32) {
	//TODO implement me
	m.Len = len
}

func (m *BaseMessage) SetData(data []byte) {
	//TODO implement me
	m.Data = data
}
