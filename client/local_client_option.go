package client

import "github.com/spf13/cast"

type LocalClientOption func(*LocalClient)

func WithPath(path *string) LocalClientOption {
	return func(client *LocalClient) {
		client.path = cast.ToString(path)
	}
}
