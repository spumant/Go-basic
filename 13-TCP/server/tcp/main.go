package main

import (
	"Go-basic/13-TCP/common"
	"encoding/json"
	"fmt"
	"net"
	"os"
)

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(2)
	}
}

func handleOneClient(conn net.Conn) {
	defer conn.Close()
	for {
		request := make([]byte, 256)
		n, err := conn.Read(request)
		if err != nil {
			fmt.Println(err)
			break
		}

		var r common.Request
		err = json.Unmarshal(request[:n], &r)
		checkError(err)

		fmt.Println(r.A, r.B)

		response := common.Response{Sum: r.A + r.B}
		bytes, err := json.Marshal(response)
		checkError(err)

		_, err = conn.Write(bytes)
		checkError(err)
	}
}

func main() {
	addr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:80")
	checkError(err)
	listener, err := net.ListenTCP("tcp4", addr)
	checkError(err)

	for {
		conn, err := listener.Accept()
		checkError(err)
		go handleOneClient(conn)
	}
}
