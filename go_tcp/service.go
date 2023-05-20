package main

import (
	"fmt"
	"net"
)

var (
	clients    = make(map[net.Addr]net.Conn)
	addCh      = make(chan net.Conn)
	delCh      = make(chan net.Addr)
	messageCh  = make(chan []byte)
	listenAddr = "localhost:8080"
)

func main() {
	fmt.Println("Server started on", listenAddr)
	listener, err := net.Listen("tcp", listenAddr)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer listener.Close()

	go broadcaster()

	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println(err)
			continue
		}

		addCh <- conn

		go handleConn(conn)
	}
}

func broadcaster() {
	for {
		select {
		case conn := <-addCh:
			clients[conn.RemoteAddr()] = conn
			fmt.Println("New client:", conn.RemoteAddr())
		case addr := <-delCh:
			delete(clients, addr)
			fmt.Println("Client disconnected:", addr)
		case msg := <-messageCh:
			for _, conn := range clients {
				_, err := conn.Write(msg)
				if err != nil {
					fmt.Println(err)
				}
			}
		}
	}
}

func handleConn(conn net.Conn) {
	defer func() {
		delCh <- conn.RemoteAddr()
		conn.Close()
	}()

	for {
		msg := make([]byte, 4096)
		n, err := conn.Read(msg)
		if err != nil {
			return
		}

		messageCh <- msg[:n]
	}
}
