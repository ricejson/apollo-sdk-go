package client

import (
	"context"
	"errors"
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

func (c *Client) IsToggleAllow(ctx context.Context, key string, conditions map[string]any) (bool, error) {
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
