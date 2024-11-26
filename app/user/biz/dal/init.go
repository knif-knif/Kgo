package dal

import (
	"github.com/knif-knif/Kgo/app/user/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/user/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
