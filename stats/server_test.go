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
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
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
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
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
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
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
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
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
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
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
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
