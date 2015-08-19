package mail

import (
	"encoding/json"
	"fmt"
	"net/smtp"
	"strings"
)

/*
send e-mail
*/

func send(user, password, host, to, subject, body, mailtype string) error {
	hp := strings.Split(host, ":")
	auth := smtp.PlainAuth("", user, password, hp[0])
	var content_type string
	if mailtype == "html" {
		content_type = "Content-Type: text/" + mailtype + "; charset=UTF-8"
	} else {
		content_type = "Content-Type: text/plain" + "; charset=UTF-8"
	}
	msg := []byte("To: " + to + "\r\nFrom: " + user + "<" + user + ">\r\nSubject: " + subject + "\r\n" + content_type + "\r\n\r\n" + body)
	send_to := strings.Split(to, ";")
	err := smtp.SendMail(host, auth, user, send_to, msg)
	return err
}

func GetBytes(data interface{}) []byte {
	b, err := json.Marshal(data)
	if err != nil {
		return []byte("nil")
	}
	return b
}

func SendMail(content interface{}) {
	user := "kingaso@163.com"
	password := "kingaso2014"
	host := "smtp.163.com:25"
	to := "shaalx@163.com"

	subject := "--Notice--"
	data := GetBytes(content)
	body := `
    <html>
    <body>
    <h3>
    ` + string(data) + `
    </h3>
    </body>
    </html>
    `
	// fmt.Println("send email")
	err := send(user, password, host, to, subject, body, "html")
	if err != nil {
		fmt.Printf("send e-mainl to %s error!:\n", to)
		fmt.Println(err)
	} else {
		fmt.Printf("send e-mainl to %s success!\n", to)
	}

}
