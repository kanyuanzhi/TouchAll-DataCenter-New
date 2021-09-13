package main

import (
	"TouchInterface/config"
	"TouchInterface/model"
	"encoding/json"
	"fmt"
	"log"
	"net"
	"strconv"
)

var limitChan = make(chan bool, 100)

func udpProcess(conn *net.UDPConn) {
	data := make([]byte, config.SERVER_RECV_LEN)
	n, udpAddr, err := conn.ReadFromUDP(data)
	if err != nil {
		return
	}
	logString := fmt.Sprintf("数据长度：%d, 客户端地址：%s, 内容：%s", n, udpAddr, data[:n])
	log.Println(logString)
	command := model.NewCommand()
	command.Duration = 3
	commandData, _ := json.Marshal(command)
	conn.WriteToUDP(commandData, udpAddr)
	//str := string(data[:n])
	log.Println(len(limitChan))
	<-limitChan
}

func main() {
	address := config.SERVER_IP + ":" + strconv.Itoa(config.SERVER_PORT)
	log.Println(address)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	conn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println(err.Error())
		return
	}
	defer conn.Close()

	for {
		limitChan <- true
		go udpProcess(conn)
	}
}
