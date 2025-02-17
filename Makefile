.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/home.proto


gen-rpc-client:
	cwgo client --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/rpc_gen -I ../idl --idl ../idl/user.proto

gen-rpc-server:
	cwgo server --type RPC --service user --module github.com/cloudwego/biz-demo/gomall/app/user --pass "-use github.com/cloudwego/biz-demo/gomall/rpc_gen/kitex_gen" -I ../../idl --idl ../../idl/user.proto