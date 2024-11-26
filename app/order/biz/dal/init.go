package dal

import (
	"github.com/knif-knif/Kgo/app/order/biz/dal/mysql"
	"github.com/knif-knif/Kgo/app/order/biz/dal/redis"
)

func Init() {
	redis.Init()
	mysql.Init()
}
