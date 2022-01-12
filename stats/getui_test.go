package stats

import (
	"context"
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiStats_GetTasks(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetTasks(context.TODO(), []string{core.TestTasks}, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetTaskGroup(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetTaskGroup(context.TODO(), core.TestGroup, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetPushCount(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetPushCount(context.TODO(), core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetPushDataByDay(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetPushDataByDay(context.TODO(), time.Now().AddDate(0, 0, -1), core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetUserDataByDay(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetUserDataByDay(context.TODO(), time.Now().AddDate(0, 0, -1), core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetOnlineUserBy24H(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiStats(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.GetOnlineUserBy24H(context.TODO(), core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}
