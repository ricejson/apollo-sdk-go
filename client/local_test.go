package client

import (
	"context"
	"github.com/stretchr/testify/assert"
	"testing"
)

func Test_LocalClient(t *testing.T) {
	testCases := []struct {
		name    string
		path    string
		wantErr error
		wantRes bool
	}{
		{
			name:    "success",
			path:    "./toggles",
			wantErr: nil,
			wantRes: true,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			client := NewLocalClient()
			err := client.Load(context.Background(), WithPath(&tc.path))
			assert.Equal(t, tc.wantErr, err)
			allow, err := client.IsToggleAllow(context.Background(), "test", "test", nil)
			assert.Equal(t, nil, err)
			assert.Equal(t, tc.wantRes, allow)
		})
	}
}
