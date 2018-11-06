package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"gofeng/ws/impl"
	"net/http"
	"time"
)

var (
	upgrader = websocket.Upgrader{
		//允许跨域
		CheckOrigin: func(r *http.Request) bool {
			return true
		},
	}
)

func wsHandler(w http.ResponseWriter, r *http.Request) {
	var (
		wsConn *websocket.Conn
		conn   *impl.Connection
		err    error
		data   []byte
	)
	//w.Write([]byte("hello"))
	//upgrader websocket
	if wsConn, err = upgrader.Upgrade(w, r, nil); err != nil {
		return
	}

	if conn, err = impl.InitConnection(wsConn); err != nil {
		goto ERR
	}
	go func() {
		var (
			err error
		)
		for {
			if err = conn.WriteMessage([]byte("heart")); err != nil {
				return
			}
			time.Sleep(1 * time.Second)
		}

	}()

	for {
		//text binary
		if data, err = conn.ReadMessage(); err != nil {
			goto ERR
		}
		if err = conn.WriteMessage(data); err != nil {
			goto ERR
		}
	}
ERR:
	conn.Close()
	fmt.Print("Error")
}

func main() {
	http.HandleFunc("/ws", wsHandler)
	http.ListenAndServe("0.0.0.0:7777", nil)
}
