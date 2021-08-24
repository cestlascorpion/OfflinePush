package stats

import (
	"fmt"
	"testing"
	"time"

	. "github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiStats_Resp2Tasks(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetTasks([]string{TestTasks}, TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetTaskGroup(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetTaskGroup(TestGroup, TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetPushCount(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetPushCount(TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetPushDataByDay(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetPushDataByDay(time.Now().AddDate(0, 0, -1), TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetUserDataByDay(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetUserDataByDay(time.Now().AddDate(0, 0, -1), TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGeTuiStats_GetOnlineUserBy24H(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	result, err := agent.GetOnlineUserBy24H(TestToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}
