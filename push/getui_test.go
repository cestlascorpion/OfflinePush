package push

import (
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
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushSingleByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushSingleByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushSingleByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushBatchByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushBatchByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushBatchByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushBatchByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushListByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
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
	result, err := agent.PushListByCid(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushListByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	taskId, err := agent.CreateMsg(create, core.TestAuthToken)
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
	result, err := agent.PushListByAlias(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushAll(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushAll(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushByTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushByTag(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_PushByFastCustomTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

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
	result, err := agent.PushByFastCustomTag(request, core.TestAuthToken)
	if err != nil {
		t.FailNow()
	}
	fmt.Println(result)
}

func TestGeTuiPush_StopPush(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
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

func TestGeTuiPush_DeleteScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
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

func TestGeTuiPush_QueryScheduleTask(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
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

func TestGeTuiPush_QueryDetail(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		http.DefaultClient)
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
