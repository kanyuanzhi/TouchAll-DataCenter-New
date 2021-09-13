package equipmentModel

type Info struct {
	BasicInfo      *BasicInfo   `json:"basic_info"`
	RunningInfo    *RunningInfo `json:"running_info"`
	ClientDuration int          `json:"client_duration"`
}

func NewInfo() *Info {
	return &Info{
		NewBasicInfo(),
		NewRunningInfo(),
		0,
	}
}

type BasicInfo struct {
	HostBasicInfo *HostBasicInfo `json:"host_basic_info"`
	CpuBasicInfo  *CpuBasicInfo  `json:"cpu_basic_info"`
	MemBasicInfo  *MemBasicInfo  `json:"mem_basic_info"`
	NetBasicInfo  *NetBasicInfo  `json:"net_basic_info"`
	DiskBasicInfo *DiskBasicInfo `json:"disk_basic_info"`
}

func NewBasicInfo() *BasicInfo {
	return &BasicInfo{
		NewHostBasicInfo(),
		NewCpuBasicInfo(),
		NewMemBasicInfo(),
		NewNetBasicInfo(),
		NewDiskBasicInfo(),
	}
}

type RunningInfo struct {
	HostRunningInfo *HostRunningInfo `json:"host_running_info"`
	CpuRunningInfo  *CpuRunningInfo  `json:"cpu_running_info"`
	MemRunningInfo  *MemRunningInfo  `json:"mem_running_info"`
	NetRunningInfo  *NetRunningInfo  `json:"net_running_info"`
	DiskRunningInfo *DiskRunningInfo `json:"disk_running_info"`
}

func NewRunningInfo() *RunningInfo {
	return &RunningInfo{
		NewHostRunningInfo(),
		NewCpuRunningInfo(),
		NewMemRunningInfo(),
		NewNetRunningInfo(),
		NewDiskRunningInfo(),
	}
}
