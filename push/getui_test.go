package push

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/configor"
)

func TestGeTuiPush_PushSingleByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &SingleReq{
		RequestId: bson.NewObjectId().Hex(),
		Audience: &proto.Audience{
			Cid: []string{core.TestToken},
		},
		Settings: nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushSingleByCid(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushSingleByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &SingleReq{
		RequestId: bson.NewObjectId().Hex(),
		Audience: &proto.Audience{
			Alias: []string{core.TestAlias},
		},
		Settings: nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushSingleByAlias(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushBatchByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &BatchReq{
		IsAsync: false,
		MsgList: []*SingleReq{
			{
				RequestId: bson.NewObjectId().Hex(),
				Audience: &proto.Audience{
					Cid: []string{core.TestToken},
				},
				Settings: nil,
				PushMessage: &proto.PushMessage{
					Notification: &proto.PushMessage_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
				PushChannel: &proto.PushChannel{
					Android: &proto.AndroidChannel{
						Ups: &proto.AndroidChannel_Ups{
							Notification: &proto.AndroidChannel_Ups_Notification{
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
	result, err := agent.PushBatchByCid(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushBatchByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &BatchReq{
		IsAsync: false,
		MsgList: []*SingleReq{
			{
				RequestId: bson.NewObjectId().Hex(),
				Audience: &proto.Audience{
					Alias: []string{core.TestAlias},
				},
				Settings: nil,
				PushMessage: &proto.PushMessage{
					Notification: &proto.PushMessage_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
				PushChannel: &proto.PushChannel{
					Android: &proto.AndroidChannel{
						Ups: &proto.AndroidChannel_Ups{
							Notification: &proto.AndroidChannel_Ups_Notification{
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
	result, err := agent.PushBatchByAlias(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushListByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	create := &CreateReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Settings:  nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	taskId, err := agent.CreateMsg(context.TODO(), create, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(taskId)

	request := &ListReq{
		Audience: &proto.Audience{
			Cid: []string{core.TestToken},
		},
		IsAsync: false,
		TaskId:  taskId,
	}
	result, err := agent.PushListByCid(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushListByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	create := &CreateReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Settings:  nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	taskId, err := agent.CreateMsg(context.TODO(), create, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(taskId)

	request := &ListReq{
		Audience: &proto.Audience{
			Alias: []string{core.TestAlias},
		},
		IsAsync: false,
		TaskId:  taskId,
	}
	result, err := agent.PushListByAlias(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushAll(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &AllReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience:  "all",
		Settings:  nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushAll(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushByTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &ByTagReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience: &proto.Audience{
			Tag: []*proto.Audience_Tag{
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
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushByTag(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushByFastCustomTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	request := &ByTagReq{
		RequestId: bson.NewObjectId().Hex(),
		GroupName: core.TestGroup,
		Audience: &proto.Audience{
			FastCustomTag: "group_manager",
		},
		Settings: nil,
		PushMessage: &proto.PushMessage{
			Notification: &proto.PushMessage_Notification{
				Title:     "title",
				Body:      "body",
				ClickType: "startapp",
			},
		},
		PushChannel: &proto.PushChannel{
			Android: &proto.AndroidChannel{
				Ups: &proto.AndroidChannel_Ups{
					Notification: &proto.AndroidChannel_Ups_Notification{
						Title:     "title",
						Body:      "body",
						ClickType: "startapp",
					},
				},
			},
		},
	}
	result, err := agent.PushByFastCustomTag(context.TODO(), request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_StopPush(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.StopPush(context.TODO(), core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_DeleteScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.DeleteScheduleTask(context.TODO(), core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_QueryScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.QueryScheduleTask(context.TODO(), core.TestTaskId, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_QueryDetail(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.QueryDetail(context.TODO(), core.TestTaskId, core.TestToken, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}
