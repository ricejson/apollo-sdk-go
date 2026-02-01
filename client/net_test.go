package client

import (
	"context"
	"fmt"
	"github.com/ricejson/apollo-sdk-go/model"
	"testing"
)

func TestNetClient_IsToggleAllow(t *testing.T) {
	client := NewNetClient()
	allow, err := client.IsToggleAllow(context.Background(), "gs_test_toggle", "", model.NewUser().With("user_id", "1"))
	fmt.Println(allow, err)
}
