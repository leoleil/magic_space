package asd

import (
	"encoding/base64"
	"crypto/md5"
	"errors"
	"fmt"
	"github.com/leoleil/magic_space/common/email"
	asdDao "github.com/leoleil/magic_space/module/user"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

// 验证
func Authentication(username, password string) (key string, err error) {
	user, err := asdDao.QueryUserByUsername(username)
	if err != nil {
		return
	}

	if user.Psw == encrypt(password) {
		onlyKey, _ := uuid.NewV4() // 生成秘钥
		key := onlyKey.String()
		err = asdDao.UpdateKeyByUserId(user.Id, key)
		return key, err
	} else {
		err = errors.New("密码或者用户名不正确")
		return key, err
	}
}

// 授权
func Authorization(key string) (user asdDao.MsSysUser, err error) {
	return asdDao.QueryUserByKey(key)
}

// 注册
func SignIn(username, password, passwordAgain, mail string) error {
	if strings.Compare(password, passwordAgain) != 0 {
		return errors.New("两次输入密码不一致")
	}
	// 判断是否邮箱已经被注册
	if emailNums, confirm, err := asdDao.QueryUserConfirmByEmail(mail); err == nil {
		if emailNums > 1 || confirm {
			return errors.New("邮箱已被注册")
		} else if emailNums == 1 && confirm {
			return errors.New("邮箱已被注册")
		}
	}
	if !email.SendToSomeConfirm(mail) {
		return errors.New("发送验证邮件失败")
	}
	err := asdDao.InsertUser(username, password, mail)

	err := asdDao.InsertUser(username, encrypt(password), mail)
	return err
}

// 加密
func encrypt(key string) string {
	data := []byte(key)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}

// 激活
func Activation(key, email string) error {

}

func sendConfirmMail(username, email string) bool {
	// 定义收件人
	mailTo := []string{email}
	// 邮件主题
	subject := "MC Space 激活邮件"
	// key 用户信息和email一起做md5加密

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
