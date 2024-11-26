package dal

import (
	"github.com/knif-knif/Kgo/app/frontend/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/frontend/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
