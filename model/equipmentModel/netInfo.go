package equipmentModel

type NetInfo struct {
	NetBasicInfo   *NetBasicInfo   `json:"net_basic_info"`
	NetRunningInfo *NetRunningInfo `json:"net_running_info"`
}

func NewNetInfo() *NetInfo {
	return &NetInfo{
		NewNetBasicInfo(),
		NewNetRunningInfo(),
	}
}

type NetBasicInfo struct {
	Name string `json:"name"`
	IP   string `json:"ip"`
	Mac  string `json:"mac"`
}

func NewNetBasicInfo() *NetBasicInfo {
	return &NetBasicInfo{}
}

type NetRunningInfo struct {
	BytesSent   uint64 `json:"bytes_sent"`   // number of bytes sent
	BytesRecv   uint64 `json:"bytes_recv"`   // number of bytes received
	PacketsSent uint64 `json:"packets_sent"` // number of packets sent
	PacketsRecv uint64 `json:"packets_recv"` // number of packets received
	Errin       uint64 `json:"errin"`        // total number of errors while receiving
	Errout      uint64 `json:"errout"`       // total number of errors while sending
	Dropin      uint64 `json:"dropin"`       // total number of incoming packets which were dropped
	Dropout     uint64 `json:"dropout"`      // total number of outgoing packets which were dropped (always 0 on OSX and BSD)
}

func NewNetRunningInfo() *NetRunningInfo {
	return &NetRunningInfo{}
}
