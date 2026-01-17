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

// Client 客户端
type Client struct {
	toggles map[string]*toggles.Toggle
}

func NewClient() *Client {
	// TODO: 读取本地配置加载所有开关
	toggles := map[string]*toggles.Toggle{}
	return &Client{
		toggles: toggles,
	}
}

// IsToggleAllow 判断开关是否允许进入
func (c *Client) IsToggleAllow(ctx context.Context, key string, userId string, conditions map[string]any) (bool, error) {
	toggle, ok := c.toggles[key]
	if !ok || toggle == nil {
		return false, ErrToggleNotFound
	}
	for _, audience := range toggle.Audiences {
		allow := true
		for _, rule := range audience.Rules {
			compareRes := rule.Compare(conditions)
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

// IsToggleAllowV2 判断开关是否允许进入
func (c *Client) IsToggleAllowV2(ctx context.Context, key string, userId string, user *model.User) (bool, error) {
	return c.IsToggleAllow(ctx, key, userId, user.Conditions)
}
