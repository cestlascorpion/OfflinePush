package user

import (
	"context"
	"fmt"
	"testing"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/jinzhu/configor"
)

func TestServer_BindAlias(t *testing.T) {
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
	resp, err := svr.BindAlias(context.Background(), &proto.BindAliasReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		DataList: []*proto.BindAliasReq_Data{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryAliasByCid(t *testing.T) {
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
	resp, err := svr.QueryAliasByCid(context.Background(), &proto.QueryAliasReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CId:       core.TestToken,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryCidByAlias(t *testing.T) {
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
	resp, err := svr.QueryCidByAlias(context.Background(), &proto.QueryCidReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		Alias:     core.TestAlias,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestMgr_UnbindAlias(t *testing.T) {
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
	resp, err := svr.UnbindAlias(context.Background(), &proto.UnbindAliasReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		DataList: []*proto.UnbindAliasReq_Data{
			{
				Cid:   core.TestToken,
				Alias: core.TestAlias,
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_RevokeAlias(t *testing.T) {
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
	resp, err := svr.RevokeAlias(context.Background(), &proto.RevokeAliasReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		Alias:     core.TestAlias,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_BindUserWithTag(t *testing.T) {
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
	resp, err := svr.BindUserWithTag(context.Background(), &proto.BindUserWithTagReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CId:       core.TestToken,
		TagList:   []string{core.TestTag},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_BindTagWithUser(t *testing.T) {
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
	resp, err := svr.BindTagWithUser(context.Background(), &proto.BindTagWithUserReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		Tag:       core.TestTag,
		CIdList:   []string{core.TestToken},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_UnbindTagFromUser(t *testing.T) {
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
	resp, err := svr.UnbindTagFromUser(context.Background(), &proto.UnbindTagFromUserReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		Tag:       core.TestTag,
		CIdList:   []string{core.TestToken},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryUserTag(t *testing.T) {
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
	resp, err := svr.QueryUserTag(context.Background(), &proto.QueryUserTagReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CId:       core.TestToken,
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryDeviceStatus(t *testing.T) {
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
	resp, err := svr.QueryDeviceStatus(context.Background(), &proto.QueryDeviceStatusReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CIdList:   []string{core.TestToken},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryUserInfo(t *testing.T) {
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
	resp, err := svr.QueryUserInfo(context.Background(), &proto.QueryUserInfoReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CIdList:   []string{core.TestToken},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_SetPushBadge(t *testing.T) {
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
	resp, err := svr.SetPushBadge(context.Background(), &proto.SetPushBadgeReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		CIdList:   []string{core.TestToken},
		Operation: "+1",
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}

func TestServer_QueryUserCount(t *testing.T) {
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
	resp, err := svr.QueryUserCount(context.Background(), &proto.QueryUserCountReq{
		PushAgent: conf.GeTui.AgentId,
		BundleId:  conf.GeTui.BundleId,
		TagList: []*proto.QueryUserCountReq_Tag{
			{
				Key:     "custom_tag",
				Values:  []string{core.TestTag},
				OptType: "and",
			},
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
