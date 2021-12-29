package core

import (
	"fmt"
	"sync"
	"testing"
	"time"

	"github.com/jinzhu/configor"
)

func TestAuthCache_Start(t *testing.T) {
	conf := &PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}

	cache, err := NewAuthCache()
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer cache.Close()

	var wg sync.WaitGroup
	wg.Add(100)
	for i := 0; i < 100; i++ {
		go func(id int) {
			defer wg.Done()
			for j := 0; j < 100; j++ {
				time.Sleep(time.Second * 10)
				uniqueId := UniqueId{PushAgent: conf.GeTui.AgentId, BundleId: conf.GeTui.BundleId}
				resp, err := cache.GetAuth(uniqueId)
				if err != nil {
					fmt.Printf("%d-%d get auth failed, %s %+v\n", id, j, uniqueId, err)
					continue
				}
				if len(resp.Token) == 0 {
					fmt.Printf("%d-%d get auth failed, %s empty token\n", id, j, uniqueId)
					continue
				}
				fmt.Printf("%d-%d get auth ok, %s %s\n", id, j, uniqueId, resp.Token)
			}
		}(i)
		time.Sleep(time.Second * 2)
	}
	wg.Wait()
}
