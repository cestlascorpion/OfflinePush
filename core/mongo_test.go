package core

import (
	"fmt"
	"testing"

	"github.com/jinzhu/configor"
)

func TestAuthDao_GetToken(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dao, err := NewAuthDao(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	uniqueId := UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}
	auth, err := dao.GetAuth(uniqueId)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(auth)
}

func TestAuthDao_SetToken(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.yml")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	dao, err := NewAuthDao(conf)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	uniqueId := UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}
	auth := &AuthToken{Token: TestAuthToken, ExpireAt: TestAuthExpireAt}
	err = dao.SetAuth(uniqueId, auth)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(auth)
}
