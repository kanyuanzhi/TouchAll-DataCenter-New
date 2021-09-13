package socket

import (
	"TouchInterface/config"
	"TouchInterface/model"
	"TouchInterface/model/equipmentModel"
	"TouchInterface/stream"
	"encoding/json"
	"log"
	"net"
)

type EquipmentServer struct {
	*Server
}

func NewEquipmentServer() *EquipmentServer {
	server := NewServer(config.EQUIPMENT_SERVER_PORT)
	return &EquipmentServer{
		server,
	}
}

func (server *EquipmentServer) SetStream(stream *stream.Stream) {
	server.stream = stream
}

func (server *EquipmentServer) Start() {
	go func() {
		defer server.udpConn.Close()
		for {
			server.limitConn <- true
			go server.receiveInfo()
		}
	}()
}

func (server *EquipmentServer) receiveInfo() {
	buf := make([]byte, config.SERVER_RECV_LEN)
	n, _, err := server.udpConn.ReadFromUDP(buf)
	if err != nil {
		log.Println(err.Error())
	}
	//logString := fmt.Sprintf("数据长度：%d, 客户端地址：%s, 内容：%s", n, clientAddr, "")
	//log.Println(logString)
	//server.sendCommand(clientAddr)

	var info equipmentModel.Info
	err = json.Unmarshal(buf[:n], &info)
	if err != nil {
		log.Println(err.Error())
		return
	}

	clientMac := info.BasicInfo.NetBasicInfo.Mac

	server.stream.EquipmentRWMutex.Lock()
	server.stream.EquipmentInfos[clientMac] = info
	server.stream.EquipmentRWMutex.Unlock()

	<-server.limitConn
}

func (server *EquipmentServer) sendCommand(clientAddr *net.UDPAddr) {
	command := model.NewCommand()
	command.Duration = 2
	commandData, _ := json.Marshal(command)
	server.udpConn.WriteToUDP(commandData, clientAddr)
}
