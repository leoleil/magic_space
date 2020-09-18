package email

import (
	"github.com/leoleil/magic_space/common/config"
	"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.AppHandle.GetConf("../../config/app.yml")
}

func TestSendToSomeConfirm(t *testing.T) {
	assert.Equal(t, true, SendToSomeConfirm("1065605635@qq.com"))
}
