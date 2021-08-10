package auth

import (
	"fmt"
	"testing"
	"time"

	. "github.com/cestlascorpion/push/core"
	"github.com/jinzhu/configor"
)

func TestGeTuiAuth_GetAuthToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiAgent(
		GTBaseUrl,
		conf.TestApp.AppId,
		conf.TestApp.AppKey,
		conf.TestApp.MasterSecret,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	auth, err := agent.GetAuth()
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(auth)
}

func TestGeTuiAuth_DelAuthToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	agent, err := NewGeTuiAgent(
		GTBaseUrl,
		conf.TestApp.AppId,
		conf.TestApp.AppKey,
		conf.TestApp.MasterSecret,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	err = agent.DelAuth(TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(TestAuthToken)
}
