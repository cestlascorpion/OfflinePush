package user

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiUser_BindAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.BindAlias(context.TODO(), &AliasList{
		DataList: []*DataList{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_QueryAliasByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryAliasByCid(context.TODO(), core.TestToken, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryCidByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryCidByAlias(context.TODO(), core.TestAlias, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_UnbindAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.UnbindAlias(context.TODO(), &AliasList{
		DataList: []*DataList{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_RevokeAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.RevokeAlias(context.TODO(), core.TestAlias, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_BindUserWithTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.BindUserWithTag(context.TODO(), core.TestToken, &CustomTagList{
		TagList: []string{core.TestTag},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_BindTagWithUser(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.BindTagWithUser(context.TODO(), core.TestTag, &CidList{
		CidList: []string{core.TestToken},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_UnbindTagFromUser(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.UnbindTagFromUser(context.TODO(), core.TestTag, &CidList{
		CidList: []string{core.TestToken},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryUserTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryUserTag(context.TODO(), core.TestToken, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_AddBlackList(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.AddBlackList(context.TODO(), []string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_DelBlackList(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.DelBlackList(context.TODO(), []string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_QueryUserStatus(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryUserStatus(context.TODO(), []string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryDeviceStatus(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryDeviceStatus(context.TODO(), []string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryUserInfo(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	invalid, validDetail, err := agent.QueryUserInfo(context.TODO(), []string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(invalid)
	fmt.Println(validDetail)
}

func TestGeTuiUser_SetPushBadge(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.SetPushBadge(context.TODO(), []string{core.TestToken}, &Operation{
		Badge: "1",
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}

func TestGeTuiUser_QueryUserCount(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	resp, err := agent.QueryUserCount(context.TODO(), &ComplexTagList{
		Tag: []*Tag{
			{
				Key:     "custom_tag",
				Values:  []string{core.TestTag},
				OptType: "and",
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
