package dal

import (
	"github.com/knif-knif/Kgo/app/payment/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
