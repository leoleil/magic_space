package asd

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/leoleil/magic_space/common/utilities"
	asdService "github.com/leoleil/magic_space/service/asd"
)

func Login(context *gin.Context) {
	username := context.DefaultPostForm("username", "leo")
	psw := context.DefaultPostForm("password", "")
	key, err := asdService.Authentication(username, psw)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "用户名或者密码不正确",
		})
		return
	}
	//utilities.SetKey(context,key)
	context.JSON(http.StatusOK, gin.H{
		"msg":      "登录成功",
		"key":      key,
		"username": username,
	})
	return
}
func SignIn(context *gin.Context) {
	username := context.DefaultPostForm("username", "leo")
	password := context.DefaultPostForm("password", "")
	passwordAgain := context.DefaultPostForm("passwordAgain", "")
	mail := context.DefaultPostForm("mail", "")
	err := asdService.SignIn(username, password, passwordAgain, mail)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg": "注册成功",
	})
	return
}
func Check(context *gin.Context) {
	key := utilities.GetKey(context)
	if key == "" {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "未登陆",
		})
	}
	user, err := asdService.Authorization(key)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": err,
		})
		return
	}
	context.JSON(http.StatusOK, gin.H{
		"msg":      "验证成功",
		"key":      key,
		"username": user.Username,
	})
	return
}
