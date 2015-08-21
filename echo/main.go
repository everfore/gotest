package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

var site string

func main() {
	http.HandleFunc("/", echo)
	http.HandleFunc("/site", siteHandFunc)
	http.ListenAndServe(":8080", nil)
}

func echo(rw http.ResponseWriter, req *http.Request) {
	URL, err := url.Parse(req.RequestURI)
	if checkerr(err) {
		return
	}
	if strings.Contains(URL.String(), "/site") {
		http.Redirect(rw, req, URL.String(), 302)
		return
	}
	fmt.Println(URL)
	realURL := fmt.Sprintf("%s%s", site, URL)
	fmt.Println(realURL)
	resp, err := http.Get(realURL)
	if checkerr(err) {
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if checkerr(err) {
		return
	}
	rw.Write(bs)
}

func siteHandFunc(rw http.ResponseWriter, req *http.Request) {
	URL, err := url.Parse(req.RequestURI)
	fmt.Println(URL)
	if checkerr(err) {
		rw.Write([]byte(err.Error()))
		return
	}
	site = URL.Query().Get("site")
	fmt.Println(site)
	if len(site) <= 1 {
		return
	}
	resp, err := http.Get(site)
	if checkerr(err) {
		return
	}
	bs, err := ioutil.ReadAll(resp.Body)
	if checkerr(err) {
		return
	}
	rw.Write(bs)
	// http.Redirect(rw, req, site, 302)
}

func checkerr(err error) bool {
	if nil != err {
		fmt.Println(err)
		return true
	}
	return false
}
