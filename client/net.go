package client

import (
	"context"
	"github.com/ricejson/apollo-idl-go/proto"
	"github.com/ricejson/apollo-sdk-go/model"
	"github.com/ricejson/apollo-sdk-go/toggles"
	"google.golang.org/grpc"
	"log"
)

type NetClient struct {
	toggles map[string]*toggles.Toggle
}

func NewNetClient() *NetClient {
	// TODO: 替换为真实 gRPC 服务地址
	conn, err := grpc.Dial("localhost:8992", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("net.Connect err: %v", err)
		return nil
	}
	client := proto.NewRPCToggleServiceClient(conn)
	resp, err := client.FindAll(context.Background(), &proto.FindAllReq{})
	if err != nil {
		log.Fatalf("client.FindAll err: %v", err)
		return nil
	}
	toggleList := make(map[string]*toggles.Toggle)
	for _, v := range resp.Toggles {
		toggleList[v.Key] = convertToggle(v)
	}
	return &NetClient{
		toggles: toggleList,
	}
}

func convertToggle(toggle *proto.Toggle) *toggles.Toggle {
	return &toggles.Toggle{
		Id:          toggle.Id,
		Name:        toggle.Name,
		Key:         toggle.Key,
		Description: toggle.Description,
		Status:      toggle.Status,
		CreateAt:    toggle.CreateAt,
		UpdateAt:    toggle.UpdateAt,
		Audiences:   convertAudiences(toggle.Audiences),
	}
}

func convertAudiences(audiences []*proto.Audience) []*toggles.Audience {
	res := make([]*toggles.Audience, 0)
	for _, audience := range audiences {
		res = append(res, &toggles.Audience{
			Id:    audience.Id,
			Name:  audience.Name,
			Rules: convertRules(audience.Rules),
		})
	}
	return res
}

func convertRules(rules []*proto.Rule) []*toggles.Rule {
	res := make([]*toggles.Rule, 0)
	for _, rule := range rules {
		res = append(res, &toggles.Rule{
			Id:        rule.Id,
			Attribute: rule.Attribute,
			Operator:  rule.Operator,
			Value:     rule.Value,
		})
	}
	return res
}

func (c *NetClient) IsToggleAllow(ctx context.Context, key string, userId string, user *model.User) (bool, error) {
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
