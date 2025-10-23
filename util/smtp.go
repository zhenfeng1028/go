package main

import (
	"log"

	"gopkg.in/gomail.v2"
)

func main() {
	// 创建新的邮件消息
	mailer := gomail.NewMessage()

	// 设置发件人、收件人、主题和正文
	from := "lizhenfeng1028@163.com" // 你的163邮箱地址
	password := "FSvws7dy5uKWihhf"   // 你的163邮箱授权码
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", "940405259@qq.com") // 收件人邮箱地址
	mailer.SetHeader("Subject", "Test email from Go using gomail")
	mailer.SetBody("text/plain", "This is the body of the email.") // 纯文本正文
	// 如果需要发送HTML内容，可以使用：
	// mailer.SetBody("text/html", "<h1>This is the body in HTML</h1>")

	// 创建SMTP拨号器
	// 注意：gomail库默认使用SSL/TLS，这里使用465端口
	dialer := gomail.NewDialer("smtp.163.com", 465, from, password)

	// 发送邮件
	if err := dialer.DialAndSend(mailer); err != nil {
		log.Fatal("Could not send email: ", err)
	}
	log.Println("Email sent successfully!")
}
