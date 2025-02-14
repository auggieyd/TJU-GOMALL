package service

import (
	"context"

	home "github.com/cloudwego/biz-demo/gomall/app/frontend/hertz_gen/frontend/home"
	"github.com/cloudwego/hertz/pkg/app"
)

type HomeService struct {
	RequestContext *app.RequestContext
	Context        context.Context
}

func NewHomeService(Context context.Context, RequestContext *app.RequestContext) *HomeService {
	return &HomeService{RequestContext: RequestContext, Context: Context}
}

func (h *HomeService) Run(req *home.Empty) (map[string]any, error) {
	var resp = make(map[string]any)
	items := []map[string]any{
		{"Name": "item1-1", "Price": 100, "Picture": "/static/image/item1.jpg"},
		{"Name": "item1-2", "Price": 110, "Picture": "/static/image/item1.jpg"},
		{"Name": "item1-3", "Price": 120, "Picture": "/static/image/item1.jpg"},
		{"Name": "item1-4", "Price": 130, "Picture": "/static/image/item1.jpg"},
		{"Name": "item2-1", "Price": 140, "Picture": "/static/image/item2.jpg"},
	}
	resp["Title"] = "Hot Sale"
	resp["Items"] = items
	return resp, nil
}
