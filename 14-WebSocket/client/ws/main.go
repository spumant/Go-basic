package main

import (
	"Go-basic/13-TCP/common"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"net/http"
)

func main() {
	dialer := &websocket.Dialer{}
	header := http.Header{
		"Name": []string{"xhz"},
	}
	conn, resp, err := dialer.Dial("ws://localhost:5657", header)
	if err != nil {
		fmt.Println("dail失败", err.Error())
		return
	}

	fmt.Println(resp.StatusCode)
	msg, _ := io.ReadAll(resp.Body)
	fmt.Println(string(msg))

	for k, v := range resp.Header {
		fmt.Println("key=", k, "value=", v)
	}

	defer conn.Close()
	for i := 0; i < 5; i++ {
		request := common.Request{A: 3, B: 5}
		conn.WriteJSON(request)
		var response common.Response
		conn.ReadJSON(&response)

		fmt.Println(response.Sum)
	}

}
