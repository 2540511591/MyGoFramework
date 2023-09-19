package iface

type IMessage interface {
	GetId() uint32
	GetLen() uint32
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
	return m.Len
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
