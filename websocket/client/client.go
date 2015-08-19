package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	// "net/http"
)

var (
	url_   = "ws://localhost:8080/echo"
	origin = "http://localhost:8080"
)

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
func main() {
	conn, err := websocket.Dial(url_, "", origin)
	checkerr(err)
	conn.Write([]byte("hello,websocket!"))
	data := make([]byte, 120)
	n, err := conn.Read(data)
	checkerr(err)
	fmt.Printf("receive msg: %s\n", data[:n])
	defer conn.Close()
}
