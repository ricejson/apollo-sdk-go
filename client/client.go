package client

import (
	"context"
	"github.com/ricejson/apollo-sdk-go/model"
)

type Client interface {
	IsToggleAllow(ctx context.Context, key string, userId string, user *model.User) (bool, error)
}
