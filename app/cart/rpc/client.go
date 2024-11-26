package rpc

import (
	"github.com/cloudwego/kitex/client"
	consul "github.com/kitex-contrib/registry-consul"
	"github.com/knif-knif/Kgo/app/cart/conf"
	cartUtils "github.com/knif-knif/Kgo/app/cart/utils"
	"github.com/knif-knif/Kgo/rpc_gen/kitex_gen/product/productcatalogservice"
	"sync"
)

var (
	ProductClient productcatalogservice.Client
	once          sync.Once
)

func InitClient() {
	var opts []client.Option
	r, err := consul.NewConsulResolver(conf.GetConf().Registry.RegistryAddress[0])
	cartUtils.MustHandleErr(err)
	opts = append(opts, client.WithResolver(r))
	ProductClient, err = productcatalogservice.NewClient("product", opts...)
	cartUtils.MustHandleErr(err)
}
