// 汉字
package main

import (
	"fmt"
	"golang.org/x/text/encoding/simplifiedchinese"
	"io/ioutil"
	"unicode/utf8"
)

func main() {
	bs, err := ioutil.ReadFile("main.md")
	if !utf8.Valid(bs) {
		fmt.Println("GBK")
		dst := make([]byte, len(bs)*2, len(bs)*3)
		simplifiedchinese.GBK.NewDecoder().Transform(dst, bs, true)
		fmt.Println(string(dst))

		dst2 := make([]byte, len(dst)/3+3)
		simplifiedchinese.GB18030.NewEncoder().Transform(dst2, dst, true)
		fmt.Println(string(dst2))
		return
	}
	fmt.Println(string(bs))
	fmt.Println(err)
}
