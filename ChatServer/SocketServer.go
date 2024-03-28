package ChatServer

import (
	"fmt"
	"log"
	"net/http"
	"sync"
	"github.com/gorilla/websocket"
)



var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool{return true},
}

var table = make(map[*websocket.Conn]bool)
var mutex sync.Mutex
var count int

/*
allows for concurent connection to happend
*/

func addConnection(connection *websocket.Conn , m *sync.Mutex){
	m.Lock()
	defer m.Unlock()
	table[connection] = true
	count++

}

func decrementCount(counter *int){
	*counter -= 1
}

func WebSocketConnection(w http.ResponseWriter , r *http.Request){
	fmt.Println("current table has : " , 1+ count ,  " people. table: " ,  table)
	conn , err := upgrader.Upgrade(w,r,nil)
	if err != nil{
	log.Println("Error upgrading user connection to websockets " , err) }
	
	go addConnection(conn , &mutex)	

	defer delete(table,conn) 
	defer conn.Close()
	defer decrementCount(&count)
	for {
            // Read message from browser
            msgType, msg, err := conn.ReadMessage()
            if err != nil {
                fmt.Println(err)
            }
	/*	
	    c , _ :=  r.Cookie("UsernameCookie")
	    forconsol := c.Value
	    fmt.Printf("%s sent: %s\n", forconsol , string(msg))
            
	    // Print the message to the console
	   /* 
            fmt.Printf("%s sent: %s\n", displayName , string(msg))
	    changestr := fmt.Sprintf("%s sent: %s\n" ,displayName , string(msg))
	    newmsg := []byte(changestr)
            */
	    // Write message back to browser
	    for conn   = range table{
		    username , _ := r.Cookie("UserNameCookie")
		    displayName := username.Value
		    changestr := fmt.Sprintf("%s sent: %s\n" ,  displayName , string(msg))
		    newmsg := []byte(changestr)
         	   if err = conn.WriteMessage(msgType, newmsg); err != nil {
			   fmt.Println("Websocket Message  error:  " ,  err)
            }}
        }
}
