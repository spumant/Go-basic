package main

import (
	"Go-basic/13-Socket/common"
	"fmt"
	"net"
	"strings"
	"time"
)

var delemeter = []byte{54, 36, 35, 34}

func readDataGram(conn net.Conn) []string {
	dataGrams := make([]string, 0)

	content := make([]byte, 1024)
	n, err := conn.Read(content)
	common.CheckError(err)
	beginIndex := 0
	var j = 0
	for i := 0; i < n; i++ {
		if content[i] == delemeter[j] {
			if j == len(delemeter) {
				dataGrams = append(dataGrams, string(content[0:i-len(delemeter)]))
				beginIndex = i + 1
			}
			j += 1
		} else {
			j = 0
		}
	}

	return dataGrams
}

func main() {
	serverAddr, err := net.ResolveTCPAddr("tcp4", "0.0.0.0:80")
	common.CheckError(err)
	listener, err := net.ListenTCP("tcp4", serverAddr)
	common.CheckError(err)
	conn, err := listener.Accept()
	checkError(err)
	fmt.Println(conn.RemoteAddr().String())

	time.Sleep(300 * time.Millisecond)

	request := make([]byte, 256)
	n, err := conn.Read(request)
	checkError(err)
	fmt.Println(string(request[:n]))

	arr := strings.Split(string(request[:n]), "|")
	//for _, v := range arr {
	//	fmt.Println(v)
	//}
	for i := 0; i < len(arr)-1; i++ {
		fmt.Println(arr[i])
	}
}
