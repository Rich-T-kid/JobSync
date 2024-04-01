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


/*

Implement a way to load the message sent bewtween two users before they joined. Doesnt have to be from database implenation just yet
mabey just store as an array of strings that are appeneded to in queue orded and then loaded in


Once this is done will bust msuta easier to remove the state from the server and encapsulate in the database 


*/
