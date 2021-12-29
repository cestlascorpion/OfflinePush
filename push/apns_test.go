package push

import (
	"fmt"
	"net/http"
	"testing"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/globalsign/mgo/bson"
	"github.com/jinzhu/configor"
	"golang.org/x/net/http2"
)

func TestApnsPush_PushSingleByCid(t *testing.T) {
	conf := &core.PushConfig{}
	err := configor.Load(conf, "conf.json")
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	agent, err := NewApnsPush(
		conf.Apns.Env,
		conf.Apns.BundleId,
		&http.Client{
			Transport: &http2.Transport{},
		},
	)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	defer agent.Close()

	result, err := agent.PushSingleByCid(&SingleReq{
		RequestId: bson.NewObjectId().Hex(),
		Audience: &proto.Audience{
			Cid: []string{core.TestToken},
		},
		Settings: &proto.Settings{
			Ttl: int64(time.Hour.Seconds()) * 2,
		},
		PushMessage: nil,
		PushChannel: &proto.PushChannel{
			Ios: &proto.IosChannel{
				Aps: &proto.IosChannel_Aps{
					Alert: &proto.IosChannel_Aps_Alert{
						Title: "tittle",
						Body:  "body",
					},
				},
				Payload: "extInfo",
			},
		},
	}, core.TestAuthToken)
	if err != nil {
		fmt.Println(err)
		t.FailNow()
	}
	fmt.Println(result)
}
