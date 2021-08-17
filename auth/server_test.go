package auth

import (
	"context"
	"fmt"
	"testing"

	. "github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/jinzhu/configor"
	"google.golang.org/grpc"
)

func TestServer_GetToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	client := proto.NewAuthClient(conn)
	resp, err := client.GetToken(context.Background(), &proto.GetTokenReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetToken2(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	client := proto.NewAuthClient(conn)
	resp, err := client.GetToken(context.Background(), &proto.GetTokenReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		OldToken:  TestAuthToken})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_SetToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	client := proto.NewAuthClient(conn)
	resp, err := client.SetToken(context.Background(), &proto.SetTokenReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Token:     TestAuthToken,
		ExpireAt:  TestAuthExpireAt})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_DelToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	conn, err := grpc.Dial(AuthServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	client := proto.NewAuthClient(conn)
	resp, err := client.DelToken(context.Background(), &proto.DelTokenReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Token:     TestAuthToken})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}
