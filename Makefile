.PHONY: gen-frontend-home
gen-frontend-home:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\home.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-auth
gen-frontend-auth:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\auth_page.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-product
gen-frontend-product:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\product_page.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-cart
gen-frontend-cart:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\cart_page.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-order
gen-frontend-order:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\order_page.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl


.PHONY: gen-frontend-checkout
gen-frontend-checkout:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\checkout_page.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-category
gen-frontend-category:
	@cd app/frontend && cwgo server --type HTTP --idl ..\\..\\idl\\frontend\\category.proto --service frontend -module github.com/knif-knif/Kgo/app/frontend -I ..\\..\\idl

.PHONY: gen-frontend-user-client
gen-frontend-user-client:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\user.proto --service user -module github.com/knif-knif/Kgo/rpc_gen -I ..\\..\\idl

.PHONY: gen-frontend-user-server
gen-frontend-user-server:
	@cd app/user && cwgo server --type RPC --idl ..\\..\\idl\\user.proto --service user -module github.com/knif-knif/Kgo/app/user --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\product.proto --service product -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/product && cwgo server --type RPC --idl ..\\..\\idl\\product.proto --service product -module github.com/knif-knif/Kgo/app/product --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-cart
gen-cart:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\cart.proto --service cart -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/cart && cwgo server --type RPC --idl ..\\..\\idl\\cart.proto --service cart -module github.com/knif-knif/Kgo/app/cart --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-payment
gen-payment:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\payment.proto --service payment -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/payment && cwgo server --type RPC --idl ..\\..\\idl\\payment.proto --service payment -module github.com/knif-knif/Kgo/app/payment --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-checkout
gen-checkout:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\checkout.proto --service checkout -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/checkout && cwgo server --type RPC --idl ..\\..\\idl\\checkout.proto --service checkout -module github.com/knif-knif/Kgo/app/checkout --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-order
gen-order:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\order.proto --service order -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/order && cwgo server --type RPC --idl ..\\..\\idl\\order.proto --service order -module github.com/knif-knif/Kgo/app/order --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl

.PHONY: gen-email
gen-email:
	@cd rpc_gen && cwgo client --type RPC --idl ..\\idl\\email.proto --service email -module github.com/knif-knif/Kgo/rpc_gen -I ..\\idl
	@cd app/email && cwgo server --type RPC --idl ..\\..\\idl\\email.proto --service email -module github.com/knif-knif/Kgo/app/email --pass "-use github.com/knif-knif/Kgo/rpc_gen/kitex_gen" -I ..\\..\\idl