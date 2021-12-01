package user

import (
	"fmt"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiUser_BindAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.BindAlias(&AliasList{
		DataList: []*DataList{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_QueryAliasByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryAliasByCid(core.TestToken, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr && err.Error() != core.InvalidTargetErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryCidByAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryCidByAlias(core.TestAlias, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr && err.Error() != core.InvalidTargetErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_UnbindAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	err = agent.UnbindAlias(&AliasList{
		DataList: []*DataList{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_RevokeAlias(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.RevokeAlias(core.TestAlias, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr && err.Error() != core.InvalidTargetErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_BindUserWithTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.BindUserWithTag(core.TestToken, &CustomTagList{
		TagList: []string{core.TestTag},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr && err.Error() != core.InvalidTargetErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_BindTagWithUser(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.BindTagWithUser(core.TestTag, &CidList{
		CidList: []string{core.TestToken},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_UnbindTagFromUser(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.UnbindTagFromUser(core.TestTag, &CidList{
		CidList: []string{core.TestToken},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryUserTag(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryUserTag(core.TestToken, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_AddBlackList(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.AddBlackList([]string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_DelBlackList(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.DelBlackList([]string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_QueryUserStatus(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryUserStatus([]string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryDeviceStatus(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryDeviceStatus([]string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}

func TestGeTuiUser_QueryUserInfo(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	invalid, validDetail, err := agent.QueryUserInfo([]string{core.TestToken}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(invalid)
	fmt.Println(validDetail)
}

func TestGeTuiUser_SetPushBadge(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.SetPushBadge([]string{core.TestToken}, &Operation{
		Badge: "1",
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
}

func TestGeTuiUser_QueryUserCount(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiUser(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	resp, err := agent.QueryUserCount(&ComplexTagList{
		Tag: []*Tag{
			{
				Key:     "custom_tag",
				Values:  []string{core.TestTag},
				OptType: "and",
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		if err.Error() != core.InvalidTokenErr {
			t.FailNow()
		} else {
			fmt.Println(err)
			return
		}
	}
	fmt.Println(resp)
}
