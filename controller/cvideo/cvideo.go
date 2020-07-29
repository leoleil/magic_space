package cvideo

import (
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"github.com/leoleil/magic_space/common/config"
	"net/http"
	"strings"
)

func GetVideoList(context *gin.Context) {
	fileInfoList, err := ioutil.ReadDir(config.AppHandle.Video.Path)
	if err != nil {
		context.JSON(http.StatusExpectationFailed, gin.H{
			"msg": "获取列表失败",
		})
		return
	}
	var videoList []string
	for i := range fileInfoList {
		if !fileInfoList[i].IsDir() && strings.Contains(fileInfoList[i].Name(), ".mp4") {
			videoList = append(videoList, fileInfoList[i].Name())
		}
	}
	context.JSON(http.StatusOK, gin.H{
		"list": videoList,
	})
	return
}
func UploadVideo(context *gin.Context) {

}
