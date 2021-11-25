package push

import (
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGetuiPush_PushSingleByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushSingleReq{
		// todo
	}
	result, err := agent.PushSingleByCid(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushSingleByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushSingleReq{
		// todo
	}
	result, err := agent.PushSingleByAlias(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushBatchByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushBatchReq{
		// todo
	}
	result, err := agent.PushBatchByCid(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushBatchByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushBatchReq{
		// todo
	}
	result, err := agent.PushBatchByAlias(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushListByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	create := &CreateMsgReq{
		// todo
	}
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(taskId)

	request := &PushListReq{
		// todo
	}
	result, err := agent.PushListByCid(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushListByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	create := &CreateMsgReq{
		// todo
	}
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(taskId)

	request := &PushListReq{
		// todo
	}
	result, err := agent.PushListByAlias(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushAll(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushAllReq{
		// todo
	}
	result, err := agent.PushAll(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushByTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushByTagReq{
		// todo
	}
	result, err := agent.PushByTag(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushByFastCustomTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	request := &PushByFastCustomTagReq{
		// todo
	}
	result, err := agent.PushByFastCustomTag(request, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_StopPush(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	result, err := agent.StopPush(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_DeleteScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	result, err := agent.DeleteScheduleTask(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_QueryScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	result, err := agent.QueryScheduleTask(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}

func TestGetuiPush_QueryDetail(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	result, err := agent.QueryDetail(core.TestTaskId, core.TestToken, core.TestAuthToken)
	if err != nil {
		t.Failed()
	}
	fmt.Println(result)
}
