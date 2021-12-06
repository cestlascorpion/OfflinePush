package push

import (
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/configor"
)

func TestGetuiPush_PushSingleByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &SingleReq{
		RequestId: bson.NewObjectId().Hex(),
		Audience: &pb.Audience{
			Cid: []string{core.TestToken},
		},
		Settings: nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushSingleByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushSingleByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &SingleReq{
		RequestId: bson.NewObjectId().Hex(),
		Audience: &pb.Audience{
			Alias: []string{core.TestAlias},
		},
		Settings: nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushSingleByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushBatchByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &BatchReq{
		IsAsync: false,
		MsgList: []*SingleReq{
			{
				RequestId: bson.NewObjectId().Hex(),
				Audience: &pb.Audience{
					Cid: []string{core.TestToken},
				},
				Settings: nil,
				PushMessage: &pb.PushMessage{
					Notification: &pb.PushMessage_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
				PushChannel: &pb.PushChannel{
					Android: &pb.AndroidChannel{
						Ups: &pb.AndroidChannel_Ups{
							Notification: &pb.AndroidChannel_Ups_Notification{
								Title:     "title",
								Body:      "body",
								ClickType: "startapp",
							},
						},
					},
				},
			},
		},
	}
	result, err := agent.PushBatchByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushBatchByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &BatchReq{
		IsAsync: false,
		MsgList: []*SingleReq{
			{
				RequestId: bson.NewObjectId().Hex(),
				Audience: &pb.Audience{
					Alias: []string{core.TestAlias},
				},
				Settings: nil,
				PushMessage: &pb.PushMessage{
					Notification: &pb.PushMessage_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
				PushChannel: &pb.PushChannel{
					Android: &pb.AndroidChannel{
						Ups: &pb.AndroidChannel_Ups{
							Notification: &pb.AndroidChannel_Ups_Notification{
								Title:     "title",
								Body:      "body",
								ClickType: "startapp",
							},
						},
					},
				},
			},
		},
	}
	result, err := agent.PushBatchByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushListByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	create := &CreateReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Settings:  nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(taskId)

	request := &ListReq{
		Audience: &pb.Audience{
			Cid: []string{core.TestToken},
		},
		IsAsync: false,
		TaskId:  taskId,
	}
	result, err := agent.PushListByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushListByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	create := &CreateReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Settings:  nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(taskId)

	request := &ListReq{
		Audience: &pb.Audience{
			Alias: []string{core.TestAlias},
		},
		IsAsync: false,
		TaskId:  taskId,
	}
	result, err := agent.PushListByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushAll(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &AllReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience:  "all",
		Settings:  nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushAll(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushByTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &ByTagReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience: &pb.Audience{
			Tag: []*pb.Audience_Tag{
				{
					Key:     "phone_type",
					Values:  []string{"android", "ios"},
					OptType: "or",
				},
				{
					Key:     "region",
					Values:  []string{"11000000"},
					OptType: "and",
				},
			},
		},
		Settings: nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushByTag(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_PushByFastCustomTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	request := &ByTagReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience: &pb.Audience{
			FastCustomTag: "group_manager",
		},
		Settings: nil,
		PushMessage: &pb.PushMessage{
			Notification: &pb.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &pb.PushChannel{
			Android: &pb.AndroidChannel{
				Ups: &pb.AndroidChannel_Ups{
					Notification: &pb.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushByFastCustomTag(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_StopPush(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, err := agent.StopPush(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_DeleteScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, err := agent.DeleteScheduleTask(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_QueryScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, err := agent.QueryScheduleTask(core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGetuiPush_QueryDetail(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	result, err := agent.QueryDetail(core.TestTaskId, core.TestToken, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}
