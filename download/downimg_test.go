package main

import (
	"github.com/qiniu/log"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"net/http"
	"os"
	"testing"
)

func DownloadImg(url_ string) {
	resp, err := http.Get(url_)
	if goutils.CheckErr(err) {
		return
	}
	log.Println(resp.Header.Get("Content-Type"))
	log.Println(resp.ContentLength)
	_ = resp
	b, err := ioutil.ReadAll(resp.Body)
	if goutils.CheckErr(err) {
		return
	}
	// log.Println(b)
	file, _ := os.OpenFile("down.png", os.O_CREATE|os.O_WRONLY, 0644)
	file.Write(b)
}

func TestDownImg(t *testing.T) {
	url_ := "http://7xku3c.com1.z0.glb.clouddn.com/golang.png"
	DownloadImg(url_)
}
