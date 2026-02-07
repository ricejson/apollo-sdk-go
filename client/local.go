package client

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/ricejson/apollo-sdk-go/model"
	"github.com/ricejson/apollo-sdk-go/toggles"
	"io/ioutil"
	"os"
)

const (
	baseDir = "/apollo/toggles/base"
)

var (
	ErrToggleNotFound = errors.New("Toggle not found")
)

// LocalClient 本地客户端
type LocalClient struct {
	path    string
	toggles map[string]*toggles.Toggle
}

func NewLocalClient() *LocalClient {
	return &LocalClient{
		path:    baseDir,
		toggles: map[string]*toggles.Toggle{},
	}
}

func (c *LocalClient) Load(ctx context.Context, LocalClientOptions ...LocalClientOption) error {
	for _, optionFunc := range LocalClientOptions {
		optionFunc(c)
	}
	res, err := loadConfigFiles(c.path)
	if err != nil {
		return err
	}
	c.toggles = res
	return nil
}

// loadConfigFiles 加载路径下的开关文件
func loadConfigFiles(path string) (map[string]*toggles.Toggle, error) {
	res := map[string]*toggles.Toggle{}
	files, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}
	for _, file := range files {
		filePath := path + string(os.PathSeparator) + file.Name()
		if file.IsDir() {
			subToggles, er := loadConfigFiles(filePath)
			if er != nil {
				return nil, err
			}
			for key, toggle := range subToggles {
				res[key] = toggle
			}
		} else {
			// 单文件
			subFile, er := os.Open(filePath)
			if er != nil {
				return nil, err
			}
			content, er := ioutil.ReadAll(subFile)
			if er != nil {
				return nil, err
			}
			var t *toggles.Toggle
			er = json.Unmarshal(content, &t)
			if er != nil {
				return nil, err
			}
			res[t.Key] = t
		}
	}
	return res, nil
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
