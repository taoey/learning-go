package mail

import (
	"crypto/tls"
	"fmt"
	"gopkg.in/gomail.v2"
	"testing"
)

func TestSendTextMail(t *testing.T) {

	user := "t741494582@163.com"
	password := "t741494582"
	host := "smtp.163.com"
	port := 25

	//构建发送器
	dialer := gomail.NewDialer(host, port, user, password)
	dialer.TLSConfig = &tls.Config{InsecureSkipVerify: true}

	//构建消息
	message := gomail.NewMessage()
	message.SetHeader("From", "t741494582@163.com")
	message.SetHeader("to", "741494582@qq.com")
	message.SetAddressHeader("Cc", "741494582@qq.com", "黄为涛11")
	message.SetHeader("Subject", "你好》")

	messageText := "www富森坎塞爱范儿就发啦三分二分拉斯坎扥啊发法令奥赛反啊<p><img src=\"cid:1.JPG height=200 width=300></p>"
	message.SetBody("text/html", messageText)

	// 设置图片ID
	h := map[string][]string{"Content-ID": {"1.JPG"}}
	message.Attach("1.JPG", gomail.SetHeader(h))

	if error := dialer.DialAndSend(message); error != nil {
		fmt.Println("发送失败")
		fmt.Println(error)
	}
}
