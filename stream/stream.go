package stream

import (
	"TouchInterface/model/equipmentModel"
	"sync"
)

type Stream struct {
	HostArriveTimeRWMutex *sync.RWMutex
	HostArriveTime        map[string]int64
	HostStatusRWMutex     *sync.RWMutex
	HostStatus            map[string]bool

	EquipmentRWMutex *sync.RWMutex
	EquipmentInfos   map[string]equipmentModel.Info
}

func NewStream() *Stream {
	return &Stream{
		new(sync.RWMutex),
		make(map[string]int64),
		new(sync.RWMutex),
		make(map[string]bool),
		new(sync.RWMutex),
		make(map[string]equipmentModel.Info),
	}
}
