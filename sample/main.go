package main

import (
	"context"
	"github.com/ricejson/apollo-sdk-go/client"
	"github.com/ricejson/apollo-sdk-go/model"
)

func main() {
	// 创建本地客户端实例
	client := client.NewLocalClient()
	// 构造condition
	user := model.NewUser().
		With("user_id", "123").
		With("city", "Beijing")
	// 获取结果
	allow, err := client.IsToggleAllow(context.Background(), "tg_wri5tl24n", "123", user)
	if err != nil {
		// 处理错误逻辑
	}
	if allow {
		// 执行业务操作
	}
}
