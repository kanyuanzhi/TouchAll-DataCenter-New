package socket

import (
	"TouchInterface/config"
	"TouchInterface/stream"
	"log"
	"net"
	"strconv"
)

type Server struct {
	udpAddr   *net.UDPAddr
	udpConn   *net.UDPConn
	limitConn chan bool
	stream    *stream.Stream
}

func NewServer(port int) *Server {
	address := config.SERVER_IP + ":" + strconv.Itoa(port)
	udpAddr, err := net.ResolveUDPAddr("udp", address)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	udpConn, err := net.ListenUDP("udp", udpAddr)
	if err != nil {
		log.Println(err.Error())
		return nil
	}
	return &Server{
		udpAddr,
		udpConn,
		make(chan bool, 100),
		nil,
	}
}
