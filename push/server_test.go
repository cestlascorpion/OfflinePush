package push

import (
	"context"
	"fmt"
	"strconv"
	"testing"
	"time"
	
	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/jinzhu/configor"
)

var (
	testAudience    *proto.Audience
	testSetting     *proto.Settings
	testPushMessage *proto.PushMessage
	testPushChannel *proto.PushChannel
)

func init() {
	testAudience = &proto.Audience{
		Cid: []string{core.TestToken},
	}
	testSetting = &proto.Settings{
		Ttl: 60 * 60 * 2 * 1000,
	}
	testPushMessage = &proto.PushMessage{
		Notification: &proto.PushMessage_Notification{
			Title: "title",
			Body:  "content",
		},
	}
	testPushChannel = &proto.PushChannel{
		Android: &proto.AndroidChannel{
			Ups: &proto.AndroidChannel_Ups{
				Notification: &proto.AndroidChannel_Ups_Notification{
					Title: "title",
					Body:  "content",
				},
			},
		},
	}
}

func TestServer_PushToSingle(t *testing.T) {
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
	resp, err := svr.PushToSingle(context.Background(), &proto.PushToSingleReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		MsgList: []*proto.SingleMsg{
			{
				RequestId:   strconv.FormatInt(time.Now().Unix(), 10),
				Audience:    testAudience,
				Settings:    testSetting,
				PushMessage: testPushMessage,
				PushChannel: testPushChannel,
			},
		},
		IsAsync: false,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_PushToList(t *testing.T) {
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
	resp1, err := svr.CreateTask(context.Background(), &proto.CreateTaskReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Msg: &proto.CreateMsg{
			RequestId:   strconv.FormatInt(time.Now().Unix(), 10),
			GroupName:   core.TestGroup,
			Settings:    testSetting,
			PushMessage: testPushMessage,
			PushChannel: testPushChannel,
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp1)
	resp2, err := svr.PushToList(context.Background(), &proto.PushToListReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Msg: &proto.ListMsg{
			Audience: testAudience,
			TaskId:   resp1.TaskId,
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp2)
}

func TestServer_PushToApp(t *testing.T) {
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
	resp, err := svr.PushToApp(context.Background(), &proto.PushToAppReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		Msg: &proto.AppMsg{
			RequestId:   strconv.FormatInt(time.Now().Unix(), 10),
			GroupName:   core.TestGroup,
			Audience:    testAudience,
			Settings:    testSetting,
			PushMessage: testPushMessage,
			PushChannel: testPushChannel,
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_StopTask(t *testing.T) {
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
	resp, err := svr.StopTask(context.Background(), &proto.StopTaskReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		TaskId:    core.TestTaskId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_RemoveTask(t *testing.T) {
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
	resp, err := svr.RemoveTask(context.Background(), &proto.RemoveTaskReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		TaskId:    core.TestTaskId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_CheckTask(t *testing.T) {
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
	resp, err := svr.CheckTask(context.Background(), &proto.CheckTaskReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		TaskId:    core.TestTaskId,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_ViewDetail(t *testing.T) {
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
	resp, err := svr.ViewDetail(context.Background(), &proto.ViewDetailReq{
		PushAgent: conf.TestApp.PushAgent,
		BundleId:  conf.TestApp.BundleId,
		TaskId:    core.TestTaskId,
		Cid:       core.TestToken,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
