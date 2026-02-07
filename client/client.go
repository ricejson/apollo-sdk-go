package client

import (
	"context"
	"github.com/ricejson/apollo-sdk-go/model"
)

type Client interface {
	// Load 加载开关
	Load(ctx context.Context) error
	// IsToggleAllow 判断开关是否开启
	IsToggleAllow(ctx context.Context, key string, userId string, user *model.User) (bool, error)
}
