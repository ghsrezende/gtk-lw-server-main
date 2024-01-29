// main.go
package main

import (
	"fmt"
	"gtk-lw-server-main/log"
	"gtk-lw-server-main/server"
	"gtk-lw-server-main/util"
	"net"
)

const (
	maxReconnectAttempts = 3
)

func main() {
	ln, err := net.Listen("tcp", util.Service)
	if err != nil {
		log.Fatal("Error listening on: ", util.Service, err)
	}

	defer ln.Close()

	log.Info("Server listening on port: ", util.ConnPort)
	fmt.Printf("Server listening on port: %s\n", util.ConnPort)

	for {
		conn, err := ln.Accept()
		if err != nil {
			log.Error("Error accepting connection: ", err)
			fmt.Printf("Error accepting connection: %v\n", err)
			continue
		}

		log.Info("New GPS Tracker connected")
		go server.HandleClient(conn)
	}
}
