package main

import (
	"fmt"
	"github.com/shaalx/goutils"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := &http.Client{}
	req, _ := http.NewRequest("GET", "http://localhost:8080", nil)
	ticker := time.NewTicker(2e9)
	for {
		select {
		case <-ticker.C:
			resp, err := c.Do(req)
			if goutils.CheckErr(err) {
				break
			}
			b, err := ioutil.ReadAll(resp.Body)
			if goutils.CheckErr(err) {
				continue
			}
			fmt.Println(goutils.ToString(b))
		}
	}
}
