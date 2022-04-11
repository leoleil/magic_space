package sasd

import (
	"errors"
	"github.com/magic_space/common/config"
	"github.com/magic_space/common/email"
	"github.com/magic_space/common/encrypt"
	asdDao "github.com/magic_space/module/user"
	uuid "github.com/satori/go.uuid"
	"log"
	"strings"
)

// Authentication 验证
func Authentication(username, password string) (key string, err error) {
	user, err := asdDao.QueryUserByUsername(username)
	if err != nil {
		return key, errors.New("无效账号")
	}
	if !user.EmailCnf {
		err = errors.New("账户未激活")
		return key, err
	}
	if user.Psw == encrypt.GetMd5Key(password) {
		onlyKey, _ := uuid.NewV4() // 生成秘钥
		key := onlyKey.String()
		err = asdDao.UpdateKeyByUserId(user.Id, key)
		return key, err
	} else {
		err = errors.New("密码或者用户名不正确")
		return key, err
	}
}

// Authorization 授权
func Authorization(key string) (user asdDao.MsSysUser, err error) {
	return asdDao.QueryUserByKey(key)
}

// SignIn 注册
func SignIn(username, password, passwordAgain, mail string) error {
	if strings.Compare(password, passwordAgain) != 0 {
		return errors.New("两次输入密码不一致")
	}
	// 判断用户名的唯一性
	if emailNums, confirm, err := asdDao.QueryUserConfirmByUser(username); err == nil {
		if emailNums != 0 {
			if !confirm {
				return errors.New(username + "已被注册，请前往邮箱激活账号")
			}
			return errors.New(username + "已被注册")
		}
	}
	// 判断是否邮箱已经被注册
	if emailNums, confirm, err := asdDao.QueryUserConfirmByEmail(mail); err == nil {
		if emailNums != 0 {
			if !confirm {
				return errors.New("邮箱" + mail + "已被注册，请前往邮箱激活账号")
			}
			return errors.New("邮箱" + mail + "已被注册")
		}
	}
	if !senMailToUser(mail) {
		return errors.New("发送验证邮件失败")
	}
	err := asdDao.InsertUser(username, encrypt.GetMd5Key(password), mail)
	return err
}

// Activation 激活
func Activation(key, email string) error {
	// 验证key
	if key != encrypt.GetMd5Key(email+"mc space") {
		return errors.New("无效链接")
	}
	_, confirmByEmail, err := asdDao.QueryUserConfirmByEmail(email)
	if err != nil {
		return errors.New("系统错误")
	}
	if confirmByEmail {
		return errors.New("重复激活")
	}
	err = asdDao.UpdateUserConfirmByEmail(email)
	if err != nil {
		return errors.New("系统错误")
	}
	return nil
}

func senMailToUser(mail string) bool {
	// 定义收件人
	mailTo := []string{mail}
	// 邮件主题
	subject := "MC Space 激活邮件"
	emailEncode := encrypt.GetMd5Key(mail + "mc space")
	url := config.AppHandle.Host.Name + "/asd/sign/confirm?key=" + emailEncode + "&email=" + mail
	body := "你好! 请点击下方网址激活邮箱：\n" + url
	err := email.SendMail(mailTo, subject, body)
	if err != nil {
		log.Println(err)
		return false
	}
	return true
}
