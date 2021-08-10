package core

import (
	"fmt"
	"testing"

	"github.com/jinzhu/configor"
)

func TestAuthDao_GetToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	dao, err := NewAuthDao(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	uniqueId := UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}
	auth, err := dao.GetToken(uniqueId)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	fmt.Println(auth)
}

func TestAuthDao_SetToken(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	dao, err := NewAuthDao(conf)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	uniqueId := UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}
	auth := &AuthToken{Token: TestAuthToken, ExpireAt: TestAuthExpireAt}
	err = dao.SetToken(uniqueId, auth)
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}

	fmt.Println(auth)
}
