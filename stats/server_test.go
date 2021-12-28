package stats

import (
	"context"
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/jinzhu/configor"
)

func TestServer_GetTasks(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetTasks(context.Background(), &proto.GetTasksReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		TaskList:  []string{core.TestTasks},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_GetTaskGroup(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetTaskGroup(context.Background(), &proto.GetTaskGroupReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		Group:     core.TestGroup,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_GetPushCount(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetPushCount(context.Background(), &proto.GetPushCountReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_GetPushDataByDay(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetPushDataByDay(context.Background(), &proto.GetPushDataByDayReq{
		PushAgent:  conf.GeTui.AgentId,
		BundleId:   conf.GeTui.BundleId,
		UnixSecond: time.Now().AddDate(0, 0, -1).Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_GetUserDataByDay(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetUserDataByDay(context.Background(), &proto.GetUserDataByDayReq{
		PushAgent:  conf.GeTui.AgentId,
		BundleId:   conf.GeTui.BundleId,
		UnixSecond: time.Now().AddDate(0, 0, -1).Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_GetOnlineUserBy24H(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer svr.Close()
	resp, err := svr.GetOnlineUserBy24H(context.Background(), &proto.GetOnlineUserBy24HReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
