package utilities

import "github.com/gin-gonic/gin"

func GetKey(context *gin.Context)string  {
	c, err := context.Request.Cookie("MCK")
	if err != nil{
		return ""
	}
	return c.Value
}
func SetKey(context *gin.Context, key string)  {
	context.SetCookie("MCK", key, 1000, "/", "localhost", false, true)
}