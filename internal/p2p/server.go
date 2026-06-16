package p2p

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"io"
	"net"
)

type Server struct {
	Address string
	Peers   []string
}

func NewServer(address string, peers []string) *Server {
	return &Server{Address: address, Peers: peers}
}

func (s *Server) Start() {
	ln, err := net.Listen("tcp", s.Address)
	if err != nil {
		fmt.Println("P2P listen error:", err)
		return
	}
	fmt.Println("P2P node listening on", s.Address)
	for {
		conn, err := ln.Accept()
		if err != nil {
			continue
		}
		go s.handleConn(conn)
	}
}

func (s *Server) handleConn(conn net.Conn) {
	defer conn.Close()
	data, err := io.ReadAll(conn)
	if err != nil {
		return
	}
	var msg Message
	gob.NewDecoder(bytes.NewReader(data)).Decode(&msg)
	fmt.Printf("Received [%s] from %s\n", msg.Command, conn.RemoteAddr())
}

func (s *Server) Send(peer string, msg Message) error {
	conn, err := net.Dial("tcp", peer)
	if err != nil {
		return err
	}
	defer conn.Close()
	return gob.NewEncoder(conn).Encode(msg)
}
