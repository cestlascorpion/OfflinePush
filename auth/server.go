package auth

import (
	"context"
	"errors"
	"net/http"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*proto.UnimplementedAuthServer
	mgr *AgentMgr
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	mgr, err := NewAgentMgr()
	if err != nil {
		log.Errorf("new auth mgr err %+v", err)
		return nil, err
	}

	agent, err := NewGeTuiAgent(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		conf.TestApp.AppKey,
		conf.TestApp.MasterSecret,
		http.DefaultClient)
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}

	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}, agent)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	return &Server{
		mgr: mgr,
	}, nil
}

func (s *Server) GetToken(ctx context.Context, in *proto.GetTokenReq) (*proto.GetTokenResp, error) {
	out := &proto.GetTokenResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.mgr.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("mgr get token err %+v", err)
		return out, err
	}

	out.Token = auth.Token
	out.ExpireAt = auth.ExpireAt

	return out, nil
}

func (s *Server) DelToken(ctx context.Context, in *proto.DelTokenReq) (*proto.DelTokenResp, error) {
	out := &proto.DelTokenResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 || len(in.Token) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.mgr.DelAuth(uniqueId, in.Token)
	if err != nil {
		log.Errorf("mgr del token err %+v", err)
		return out, err
	}

	return out, nil
}

func (s *Server) Close() {
	// nothing
}
