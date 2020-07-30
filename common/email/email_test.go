package email

import (
	"github.com/leoleil/magic_space/common/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init()  {
	config.AppHandle.GetConf("../../config/app.yml")
}

func TestSendToSome(t *testing.T) {
	assert.Equal(t, true, SendToSome("1105263198@qq.com"))
}
