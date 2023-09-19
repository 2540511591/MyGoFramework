package conf

type SConfig struct {
	//协程池内工作协程的数量
	WorkerNumber uint32
	//任务队列长度
	WorkerQueueLen uint32

	//传输层协议数据包大小，tcp/udp专用,自定义解封包可无视
	DataSize uint32
}

var (
	ServerConfig *SConfig
)

func init() {
	ServerConfig = GetDefaultConfig()
}

func GetDefaultConfig() *SConfig {
	return &SConfig{
		WorkerNumber:   10,
		WorkerQueueLen: 100,
		DataSize:       512,
	}
}
