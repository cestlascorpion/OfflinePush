package stats

import (
	"context"
	"fmt"
	"testing"
	"time"

	. "github.com/cestlascorpion/push/core"
	"github.com/cestlascorpion/push/proto"
	"github.com/jinzhu/configor"
)

func TestServer_GetTasks(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetTasks(context.Background(), &proto.GetTasksReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		TaskList:  []string{TestTasks},
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetTaskGroup(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetTaskGroup(context.Background(), &proto.GetTaskGroupReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Group:     TestGroup,
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetPushCount(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetPushCount(context.Background(), &proto.GetPushCountReq{
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
		UnixSecond: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetPushDataByDay(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetPushDataByDay(context.Background(), &proto.GetPushDataByDayReq{
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
		UnixSecond: time.Now().AddDate(0, 0, -1).Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetUserDataByDay(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetUserDataByDay(context.Background(), &proto.GetUserDataByDayReq{
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
		UnixSecond: time.Now().AddDate(0, 0, -1).Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}

func TestServer_GetOnlineUserBy24H(t *testing.T) {
	conf := &StatsConfig{}
	err := configor.Load(conf, "stats.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	svr, err := NewServer(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	defer svr.Close()
	resp, err := svr.GetOnlineUserBy24H(context.Background(), &proto.GetOnlineUserBy24HReq{
		PushAgent:  conf.TestApp.PushAgent,
		BundleId:   conf.TestApp.BundleId,
		UnixSecond: time.Now().Unix(),
	})
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(resp)
}
