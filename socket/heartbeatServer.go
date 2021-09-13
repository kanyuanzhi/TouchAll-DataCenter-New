package socket

import (
	"TouchInterface/config"
	"TouchInterface/stream"
	"log"
	"time"
)

type HeartbeatServer struct {
	*Server
	duration int
	t        *time.Timer
}

func NewHeartbeatServer() *HeartbeatServer {
	server := NewServer(config.HEARTBEAT_SERVER_PORT)
	return &HeartbeatServer{
		server,
		config.HEARTBEAT_DURATION,
		time.NewTimer(time.Duration(config.HEARTBEAT_DURATION) * time.Second),
	}
}

func (server *HeartbeatServer) SetStream(stream *stream.Stream) {
	server.stream = stream
}

func (server *HeartbeatServer) Start() {
	go func() {
		defer server.udpConn.Close()
		for {
			server.limitConn <- true
			go server.receiveHeartbeat()
		}
	}()

	go func() {
		for {
			<-server.t.C
			go server.checkHostList()
			server.t.Reset(time.Duration(server.duration) * time.Second)
		}
	}()
}

func (server *HeartbeatServer) receiveHeartbeat() {
	buf := make([]byte, config.HEARTBEAT_RECV_LEN)
	n, clientAddr, err := server.udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err.Error())
	}
	heartbeatData := string(buf[:n])
	head := heartbeatData[:len(config.HEARTBEAT_HEAD)]
	if head == config.HEARTBEAT_HEAD {
		mac := heartbeatData[len(config.HEARTBEAT_HEAD):]
		arriveTime := time.Now().Unix()

		server.stream.HostArriveTimeRWMutex.Lock()
		server.stream.HostArriveTime[mac] = arriveTime
		server.stream.HostArriveTimeRWMutex.Unlock()

		server.stream.HostStatusRWMutex.Lock()
		server.stream.HostStatus[mac] = true
		server.stream.HostStatusRWMutex.Unlock()

		heartbeatResponse := config.HEARTBEAT_HEAD + "ok"
		_, err = server.udpConn.WriteToUDP([]byte(heartbeatResponse), clientAddr)
		if err != nil {
			log.Println(err.Error())
		}
	} else {
		log.Println("This is not a heartbeat packet!")
	}
	<-server.limitConn
}

func (server *HeartbeatServer) checkHostList() {
	checkTime := time.Now().Unix()
	for clientMac, arriveTime := range server.stream.HostArriveTime {
		if checkTime-arriveTime > config.HEARTBEAT_DURATION {
			server.stream.HostStatusRWMutex.Lock()
			server.stream.HostStatus[clientMac] = false
			server.stream.HostStatusRWMutex.Unlock()

			server.stream.EquipmentRWMutex.Lock()
			delete(server.stream.EquipmentInfos, clientMac)
			server.stream.EquipmentRWMutex.Unlock()
		}
	}
}
