package auth

import (
	"fmt"
	"testing"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/jinzhu/configor"
)

func TestApnsAuth_GetAuth(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewApnsAuth(conf.Apns.Key, conf.Apns.KeyId, conf.Apns.TeamId)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	token, err := agent.GetAuth()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(token)
}

func TestApnsAuth_DelAuth(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewApnsAuth(conf.Apns.Key, conf.Apns.KeyId, conf.Apns.TeamId)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	token, err := agent.GetAuth()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(token)

	err = agent.DelAuth(token.Token)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
}
