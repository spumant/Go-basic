package main

import (
	"Go-basic/13-Socket/common"
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

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp4", "127.0.0.1:80")
	checkError(err)
	conn, err := net.DialTCP("tcp4", nil, serverAddr)
	checkError(err)

	for i := 0; i < 3; i++ {
		request := common.Request{A: 3, B: 5}
		bytes, err := json.Marshal(request)
		checkError(err)

		_, err = conn.Write(bytes)
		checkError(err)

		response := make([]byte, 256)
		n, err := conn.Read(response)
		checkError(err)

		var r common.Response
		err = json.Unmarshal(response[:n], &r)
		checkError(err)
		fmt.Println(r.Sum)
	}


	conn.Close()
}
