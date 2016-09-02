package main

import (
	"github.com/qiniu/log"
	. "github.com/toukii/goutils"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", root)
	http.ListenAndServe(":8080", nil)
}
func root(rw http.ResponseWriter, req *http.Request) {
	notify := rw.(http.CloseNotifier).CloseNotify()
	go func() {
		select {
		case <-notify:
			log.Println("closed>", req.RemoteAddr)
		}
	}()
	log.Info(req.RemoteAddr)
	// ticker := time.NewTicker(2e9)
	// for {
	// 	select {
	// 	case <-ticker.C:
	// go hello(rw)
	rw.Write(ToByte("hello."))
	// 		log.Info("hello")
	// 	default:
	// 	}
	// }
}

func hello(rw http.ResponseWriter) {
	ticker := time.NewTicker(2e9)
	for {
		select {
		case <-ticker.C:
			rw.Write(ToByte("hello."))
			log.Info("hello")
		default:
		}
	}
}
