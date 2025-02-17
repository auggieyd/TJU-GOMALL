.PHONY: gen-frontend
gen-frontend:
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --service frontend -module github.com/cloudwego/biz-demo/gomall/app/frontend --idl ../../idl/frontend/home.proto