package impl

import (
	"github.com/gorilla/websocket"
	"sync"
	"errors"
)

type Connection struct {
	wsConn *websocket.Conn
	inChan chan []byte
	outChan chan []byte
	closeChan chan byte
	mutex sync.Mutex
	isClosed bool
}

func InitConnection(wsConn *websocket.Conn)(conn *Connection,err error){
   conn  = &Connection {
   	 wsConn : wsConn,
     inChan : make(chan []byte ,1000),
   	 outChan : make(chan []byte,1000),
   	 closeChan : make(chan byte,1),
   }
   //启动读协程
   go conn.readLoop()
   //启动写携程
   go conn.WriteLoop()
   return
}

func (conn *Connection) ReadMessage()(data []byte, err error){
   select {
   case data = <-conn.inChan:
   case <-conn.closeChan:
	 	err = errors.New("connection is closed")
   }
	return
}

func (conn *Connection) WriteMessage(data []byte)(err error){
	select {
	case conn.outChan <- data:
	case <-conn.closeChan:
		err = errors.New("connection is closed")
	}
	return
}

func (conn *Connection) Close(){
	//线程安全，可重入的close
	conn.wsConn.Close()
	//保证执行一次
	conn.mutex.Lock()
	if !conn.isClosed{
		close(conn.closeChan)
		conn.isClosed = true
	}
	conn.mutex.Unlock()
}

//内部实现
func (conn *Connection) readLoop(){
	var (
		data []byte
		err error
	)
	for{
	   if _,data,err = conn.wsConn.ReadMessage();err!=nil{
	   	  goto ERR
	   }
	   select {
	       //消息超过长度 会阻塞 inChan有空闲的空间
		  case conn.inChan <- data:
		  case <-conn.closeChan:
		  	//closeChan关闭的时候
			  goto ERR

	   }
	}
ERR:
	conn.Close()
}

func (conn *Connection)WriteLoop(){
	var (
		data []byte
		err error
	)
	for{
		select {
			case data =<-conn.outChan:
			case <-conn.closeChan:
				goto ERR

		}
		if err = conn.wsConn.WriteMessage(websocket.TextMessage,data); err != nil{
			goto ERR
		}

	}
ERR:
	conn.Close()
}