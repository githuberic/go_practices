package tests

import (
	"go_practices/go_gorm/exec_01/dal"
	"testing"
)

func TestAdd(t *testing.T) {
	user := new(dal.User)
	user.Name = "tome"
	user.Address = "杭州"
	user.Mobile = "13588827425"
	user.Add()          //user.id = 1，添加之后user中的id会变成数据库中生成的值
}
