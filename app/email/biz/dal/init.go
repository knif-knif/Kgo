package dal

import (
	"github.com/knif-knif/Kgo/app/email/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/email/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
