package main

import (
	"bufio"
	"fmt"
	"io"
	"net/http"
	"os"

	// "bytes"
	"time"
)

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}

func main() {
	go ticker()
	file, err := os.OpenFile("master.zip", os.O_CREATE|os.O_WRONLY, 0644)
	if checkerr(err) {
		return
	}
	ch := get("https://codeload.github.com/shaalx/sstruct/zip/master", file)
	// ch := get("https://codeload.github.com/everfore/push/zip/master", file)
	fmt.Println(<-ch)
}

func ticker() {
	cker := time.NewTicker(5e8)
	for {
		<-cker.C
		fmt.Print(".")
	}
}

func get(url_ string, file *os.File) chan bool {
	ch := make(chan bool)
	resp, err := http.Get(url_)
	if checkerr(err) {
		ch <- false
		return ch
	}
	buf := make([]byte, 1024)
	go func() {
		defer func() {
			file.Close()
			ch <- true
		}()
		rd := bufio.NewReader(resp.Body)
		for {
			n, err := rd.Read(buf)
			// fmt.Print(n, "-")
			m, werr := file.Write(buf[:n])
			if checkerr(werr) {
				break
			}
			fmt.Print(m, " ")
			if err == io.EOF {
				break
			}

		}
		// io.Copy(file, resp.Body)
	}()
	return ch
}
