package proxy

import (
	"context"
	"errors"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type MsgConverter interface {
	ToUnicast(interface{}) ([]*proto.PushToSingleReq, error)
	ToMulticast(interface{}) (*proto.CreateTaskReq, []*proto.PushToListReq, error)
	ToBroadcast(interface{}) (*proto.PushToAppReq, error)
}

type PushProxy struct {
	conv   MsgConverter
	client proto.PushClient
}

func NewPushProxy(msgConv MsgConverter) (*PushProxy, error) {
	if msgConv == nil {
		log.Error("nil converter")
		return nil, errors.New("invalid parameter")
	}
	conn, err := grpc.Dial(core.PushServerAddr, grpc.WithBlock(), grpc.WithInsecure())
	if err != nil {
		log.Errorf("grpc dial err %+v", err)
		return nil, err
	}
	return &PushProxy{
		conv:   msgConv,
		client: proto.NewPushClient(conn),
	}, nil
}

func (p *PushProxy) Unicast(ctx context.Context, msg interface{}) error {
	reqList, err := p.conv.ToUnicast(msg)
	if err != nil {
		log.Errorf("conv msg err %+v", err)
		return err
	}
	for _, req := range reqList {
		resp, err := p.client.PushToSingle(ctx, req)
		if err != nil {
			log.Errorf("push single err %+v", err)
			return err
		}
		log.Debugf("push single resp %+v", resp)
	}
	return nil
}

func (p *PushProxy) Multicast(ctx context.Context, msg interface{}) error {
	req1, req2List, err := p.conv.ToMulticast(msg)
	if err != nil {
		log.Errorf("conv msg err %+v", err)
		return err
	}
	resp1, err := p.client.CreateTask(ctx, req1)
	if err != nil {
		log.Errorf("push list err %+v", err)
		return err
	}
	for _, req2 := range req2List {
		req2.Msg.TaskId = resp1.TaskId
		resp2, err := p.client.PushToList(ctx, req2)
		if err != nil {
			log.Errorf("push list err %+v", err)
			return err
		}
		log.Debugf("push list resp %+v", resp2)
	}
	return nil
}

func (p *PushProxy) Broadcast(ctx context.Context, msg interface{}) error {
	req, err := p.conv.ToBroadcast(msg)
	if err != nil {
		log.Errorf("conv msg err %+v", err)
		return err
	}
	resp, err := p.client.PushToApp(ctx, req)
	if err != nil {
		log.Errorf("push app err %+v", err)
		return err
	}
	log.Debugf("push app resp %+v", resp)
	return nil
}
