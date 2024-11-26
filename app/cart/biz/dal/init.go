package dal

import (
	"github.com/knif-knif/Kgo/app/cart/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
