package user

import (
	"github.com/magic_space/common/config"
	"github.com/stretchr/testify/assert"

	//"github.com/stretchr/testify/assert"
	"testing"
)

func init() {
	config.AppHandle.GetConf("../../config/app.yml")
}

func TestQueryUserConfirmByEmail(t *testing.T) {
	nums, confirm, _ := QueryUserConfirmByEmail("1105263198@qq.com")
	assert.Equal(t, 1, nums)
	assert.Equal(t, true, confirm)
}

func TestUpdateUserConfirmByEmail(t *testing.T) {
	assert.Nil(t, UpdateUserConfirmByEmail("1105263198@qq.com"))
}

func TestQueryUserConfirmByUser(t *testing.T) {
	nums, confirm, _ := QueryUserConfirmByUser("xyp123")
	assert.Equal(t, 1, nums)
	assert.Equal(t, true, confirm)
}
