package proxy

import (
	"fmt"
	"testing"

	"github.com/cestlascorpion/offlinepush/proto"
)

type EmptyMsgConv struct {
	// empty
}

func (e *EmptyMsgConv) ToUnicast(interface{}) ([]*proto.PushToSingleReq, error) {
	return []*proto.PushToSingleReq{
		{
			PushAgent: "AGENT",
			BundleId:  "com.test.it",
			MsgList: []*proto.SingleMsg{
				{
					// implementation
				},
			},
		},
	}, nil
}

func (e *EmptyMsgConv) ToMulticast(interface{}) (*proto.CreateTaskReq, []*proto.PushToListReq, error) {
	return &proto.CreateTaskReq{
			PushAgent: "AGENT",
			BundleId:  "com.test.it",
			Msg:       &proto.CreateMsg{
				// implementation
			},
		}, []*proto.PushToListReq{
			{
				PushAgent: "AGENT",
				BundleId:  "com.test.it",
				Msg:       &proto.ListMsg{
					// implementation
				},
			},
		}, nil
}

func (e *EmptyMsgConv) ToBroadcast(interface{}) (*proto.PushToAppReq, error) {
	return &proto.PushToAppReq{
		PushAgent: "AGENT",
		BundleId:  "com.test.it",
		Msg:       &proto.AppMsg{
			// implementation
		},
	}, nil
}

func TestPushProxy(t *testing.T) {
	proxy, err := NewPushProxy(&EmptyMsgConv{})
	if err != nil {
		t.FailNow()
	}
	fmt.Println(proxy)
}
