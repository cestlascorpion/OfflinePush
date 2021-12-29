package core

import (
	"fmt"
	"testing"

	"github.com/jinzhu/configor"
)

func TestConfig(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	fmt.Println(conf.Mongo.Name)
	fmt.Println(conf.Mongo.Url)
	fmt.Println(conf.Mongo.DataBase)
	fmt.Println(conf.Mongo.AuthCollection)
	fmt.Println(conf.Mongo.StatsCollection)
	fmt.Println(conf.Mongo.PoolSize)
	fmt.Println(conf.GeTui.AgentId)
	fmt.Println(conf.GeTui.BundleId)
	fmt.Println(conf.GeTui.AppId)
	fmt.Println(conf.GeTui.AppKey)
	fmt.Println(conf.GeTui.MasterSecret)
	fmt.Println(conf.Apns.AgentId)
	fmt.Println(conf.Apns.BundleId)
	fmt.Println(conf.Apns.Env)
	fmt.Println(conf.Apns.Key)
	fmt.Println(conf.Apns.KeyId)
	fmt.Println(conf.Apns.TeamId)
}
