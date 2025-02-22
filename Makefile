.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/product_page.proto

	
.PHONY: gen-product
gen-product:
	@cd rpc_gen && cwgo client --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/rpc_gen -I ../idl --idl ../idl/product.proto
	@cd app/product && cwgo server --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/app/product --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/product.proto