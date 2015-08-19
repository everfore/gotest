package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
)

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
func webHandler(conn *websocket.Conn) {
	data := make([]byte, 120)
	n, err := conn.Read(data)
	checkerr(err)
	fmt.Printf("msg<------[%s]\n", string(data[:n]))
	m, err := conn.Write(data[:n])
	checkerr(err)
	fmt.Printf("msg------->[%s]\n", string(data[:m]))
	defer conn.Close()
}

func main() {
	http.Handle("/echo", websocket.Handler(webHandler))
	http.Handle("/", http.FileServer(http.Dir(".")))
	http.ListenAndServe(":8080", nil)
}
