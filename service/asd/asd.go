package asd

import (
	"crypto/md5"
	"errors"
	"fmt"
	uuid "github.com/satori/go.uuid"
	asdDao "magic_space/module/user"
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

	err := asdDao.InsertUser(username, encrypt(password), mail)
	return err
}

// 加密
func encrypt(key string) string {
	data := []byte(key)
	has := md5.Sum(data)
	return fmt.Sprintf("%x", has)
}
