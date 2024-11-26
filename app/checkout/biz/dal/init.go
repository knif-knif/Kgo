package dal

import (
	"github.com/knif-knif/Kgo/app/checkout/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/checkout/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
