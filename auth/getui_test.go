package auth

import (
	"context"
	"fmt"
	"net/http"
	"testing"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiAuth_GetAuthToken(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiAgent(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		conf.GeTui.AppKey,
		conf.GeTui.MasterSecret,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	auth, err := agent.GetAuth(context.TODO())
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(auth)
}

func TestGeTuiAuth_DelAuthToken(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewGeTuiAgent(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		conf.GeTui.AppKey,
		conf.GeTui.MasterSecret,
		http.DefaultClient)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	err = agent.DelAuth(context.TODO(), core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(core.TestAuthToken)
}
