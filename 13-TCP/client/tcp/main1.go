package main

import (
	"Go-basic/13-Socket/common"
	"fmt"
	"net"
)

var delemeter = []byte{54, 36, 35, 34}

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:80")
	common.CheckError(err)
	conn, err := net.DialTCP("tcp4", nil, serverAddr)
	common.CheckError(err)
	fmt.Println(conn.RemoteAddr().String())

	conn.Write([]byte("Hello"))
	conn.Write(delemeter)
	conn.Write([]byte("golang"))
	conn.Write(delemeter)

	conn.Close()
}
