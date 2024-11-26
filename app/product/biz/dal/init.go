package dal

import (
	"github.com/knif-knif/Kgo/app/product/biz/dal/mysql"
)

func Init() {
	//redis.Init()
	mysql.Init()
}
