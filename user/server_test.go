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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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
	err := configor.Load(conf, "conf.json")
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

func TestServer_ManageCidAndDeviceToken(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
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
	resp, err := svr.ManageCidAndDeviceToken(context.Background(), &proto.ManageCidAndDeviceTokenReq{
		PushAgent:    conf.GeTui.AgentId,
		BundleId:     conf.GeTui.BundleId,
		Manufacturer: "wx",
		DtList: map[string]string{
			"f5fc768b8f3d0ac24d6b7df65e0221b0": "oX-Yr4ICsWkJaydB8mdRD3DMahkA",
			"9497869031e1fa74adaa84458f73cc62": "oX-BfgICsWkJurdB3odRDcDjyskB",
		},
	})
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(resp)
}
