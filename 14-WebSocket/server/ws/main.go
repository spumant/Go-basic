package main

import (
	"Go-basic/13-TCP/common"
	"fmt"
	"github.com/gorilla/websocket"
	"net"
	"net/http"
	"strconv"
	"time"
)

type WsServer struct {
	listener net.Listener
	addr     string
	upgrade  *websocket.Upgrader
}

func NewWsServer(port int) *WsServer {
	ws := new(WsServer)
	ws.addr = "0.0.0.0:" + strconv.Itoa(port)
	ws.upgrade = &websocket.Upgrader{
		HandshakeTimeout: 5 + time.Second,
		ReadBufferSize:   4096,
		WriteBufferSize:  4096,
	}
	return ws
}

func (ws *WsServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	conn, err := ws.upgrade.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("升级失败", err)
		return
	}
	ws.handleOneConnection(conn)
}

func (ws *WsServer) handleOneConnection(conn *websocket.Conn) {
	defer func() {
		conn.Close()
	}()

	for { //长连接
		conn.SetReadDeadline(time.Now().Add(20 * time.Second))
		var request common.Request
		if err := conn.ReadJSON(&request); err != nil {
			if netError, ok := err.(net.Error); ok {
				if netError.Timeout() {
					fmt.Println("发生了读超时")
					return
				}
			}
			fmt.Println(err)
			return
		}
		response := common.Response{Sum: request.A + request.B}
		if err := conn.WriteJSON(response); err != nil {
			fmt.Println(err)
		}
	}
}

func (ws *WsServer) Start() error {
	//var err error
	//
	//ws.listener, err = net.Listen("tcp", ws.addr)
	if err := http.ListenAndServe(ws.addr, ws); err != nil {
		fmt.Println("失败", err.Error())
		return err
	} else {
		return nil
	}
}

func main() {
	ws := NewWsServer(5657)
	ws.Start()
}
