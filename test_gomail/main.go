package main

import (
	"gopkg.in/gomail.v2"
	"time"
)
func main() {
	m := gomail.NewMessage()
	m.SetHeader("From", "uin@qq.com")
	m.SetHeader("To", "uin@qq.com")
	//m.SetAddressHeader("Cc", "dan@example.com", "Dan")
	m.SetHeader("Subject", "Hello SMTP")
	m.SetBody("text/plain", "Hello")
	//m.Attach("/home/Alex/lolcat.jpg")

	d := gomail.NewDialer("smtp.qq.com", 587, "uin", "")

	// Send the email to Bob, Cora and Dan.
	if err := d.DialAndSend(m); err != nil {
		panic(err)
	}

	time.Sleep(time.Second * 1)
}
