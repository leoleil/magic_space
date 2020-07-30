package email

import (
	"encoding/base64"
	"fmt"
	"github.com/leoleil/magic_space/common/config"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func sendMail(mailTo []string, subject string, body string) error {
	email_info := config.AppHandle.Email
	mailConn := map[string]string{
		"user": email_info.User,
		"pass": email_info.Pwd,
		"host": email_info.Host,
		"port": email_info.Port,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "magic_space官方"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

func SendToSome(user string) bool {
	// 定义收件人
	mailTo := []string{user}
	// 邮件主题
	subject := "Hello by golang gomail from exmail.qq.com"
	//todo 邮件正文修改为验证连接
	emailEncode := base64.StdEncoding.EncodeToString([]byte(user))
	url := "http://www.mcspace.icu:4010/asd/sign/confirm?user=" + emailEncode + "&email=" + user
	fmt.Println(url)
	body := "Hello,by gomail sent\n" + url
	err := sendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return false
	}
	fmt.Println("send successfully")
	return true
}
