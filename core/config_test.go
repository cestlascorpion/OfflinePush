package core

import (
	"fmt"
	"testing"

	"github.com/jinzhu/configor"
)

func TestConfig(t *testing.T) {
	conf := &AuthConfig{}
	err := configor.Load(conf, "auth.yml")
	if err != nil {
		fmt.Println(err)
		t.Failed()
	}
	fmt.Println(conf.Name)
	fmt.Println(conf.Mongo.Name)
	fmt.Println(conf.Mongo.Url)
	fmt.Println(conf.Mongo.DataBase)
	fmt.Println(conf.Mongo.Collection)
	fmt.Println(conf.Mongo.PoolSize)
	fmt.Println(conf.TestApp.PushAgent)
	fmt.Println(conf.TestApp.BundleId)
	fmt.Println(conf.TestApp.AppId)
	fmt.Println(conf.TestApp.AppKey)
	fmt.Println(conf.TestApp.MasterSecret)
	fmt.Println(conf.TestApp.TimeoutSec)
}
