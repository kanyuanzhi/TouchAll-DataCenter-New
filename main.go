package main

import (
	"TouchInterface/socket"
	"TouchInterface/stream"
	"log"
	"time"
)

func main() {
	stream := stream.NewStream()

	heartbeatServer := socket.NewHeartbeatServer()
	heartbeatServer.SetStream(stream)
	heartbeatServer.Start()

	equipmentServer := socket.NewEquipmentServer()
	equipmentServer.SetStream(stream)
	equipmentServer.Start()

	t := time.NewTimer(time.Second)
	for {
		<-t.C
		//time.Sleep(time.Hour * 24)
		log.Println(stream.HostStatus)
		t.Reset(time.Second)
	}
}
