package client

import (
	"context"
	"errors"
	"github.com/ricejson/apollo-sdk-go/model"
	"github.com/ricejson/apollo-sdk-go/toggles"
)

var (
	ErrToggleNotFound = errors.New("Toggle not found")
)

// LocalClient 本地客户端
type LocalClient struct {
	toggles map[string]*toggles.Toggle
}

func NewLocalClient() *LocalClient {
	// TODO: 读取本地配置加载所有开关
	toggles := map[string]*toggles.Toggle{}
	return &LocalClient{
		toggles: toggles,
	}
}

// IsToggleAllow 判断开关是否允许进入
func (c *LocalClient) IsToggleAllow(ctx context.Context, key string, userId string, user *model.User) (bool, error) {
	toggle, ok := c.toggles[key]
	if !ok || toggle == nil {
		return false, ErrToggleNotFound
	}
	for _, audience := range toggle.Audiences {
		allow := true
		for _, rule := range audience.Rules {
			compareRes := rule.Compare(user.Conditions)
			if !compareRes {
				allow = false
				break
			}
		}
		if allow {
			return true, nil
		}
	}
	return false, nil
}
