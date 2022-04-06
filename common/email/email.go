package email

import (
	"encoding/base64"
	"fmt"
	"github.com/magic_space/common/config"
	"gopkg.in/gomail.v2"
	"log"
	"strconv"
)

func SendMail(mailTo []string, subject string, body string) error {
	emailInfo := config.AppHandle.Email
	mailConn := map[string]string{
		"user": emailInfo.User,
		"pass": emailInfo.Pwd,
		"host": emailInfo.Host,
		"port": emailInfo.Port,
	}
	port, _ := strconv.Atoi(mailConn["port"]) //转换端口类型为int
	m := gomail.NewMessage()
	// 处理被当作垃圾邮件的方法
	m.SetHeader("X-Mailer", "Microsoft Outlook Express 6.00.2900.2869")
	m.SetHeader("X-Priority", "3")
	m.SetHeader("X-MimeOLE", "Produced By Microsoft MimeOLE V6.00.2900.2869")
	m.SetHeader("ReturnReceipt", "1")
	m.SetHeader("X-Mailer", "Normal")
	// 设置基础信息
	m.SetHeader("From", m.FormatAddress(mailConn["user"], "magic_space官方"))
	m.SetHeader("To", mailTo...)    //发送给多个用户
	m.SetHeader("Subject", subject) //设置邮件主题
	m.SetBody("text/html", body)    //设置邮件正
	d := gomail.NewDialer(mailConn["host"], port, mailConn["user"], mailConn["pass"])
	err := d.DialAndSend(m)
	return err
}

func SendToSomeConfirm(user string) bool {
	// 定义收件人
	mailTo := []string{user}
	// 邮件主题
	subject := "MC Space 激活邮件"
	// todo 邮件正文修改为验证连接
	emailEncode := base64.StdEncoding.EncodeToString([]byte(user))
	url := "http://www.mcspace.icu:4010/asd/sign/confirm?user=" + emailEncode + "&email=" + user
	fmt.Println(url)
	// todo 修改正文发送格式
	body := "你好，" + user + "，请点击下方网址激活邮箱：" + url
	err := SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		fmt.Println("send fail")
		return false
	}
	fmt.Println("send successfully")
	return true
}
