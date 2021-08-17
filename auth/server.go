package auth

import (
	"context"
	"errors"
	"time"

	. "github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Dao *AuthDao
	Mgr *AgentMgr
}

func NewServer(conf *AuthConfig) (*Server, error) {
	dao, err := NewAuthDao(conf)
	if err != nil {
		log.Errorf("new auth dao err %+v", err)
		return nil, err
	}

	mgr, err := NewAgentMgr()
	if err != nil {
		log.Errorf("new auth mgr err %+v", err)
		return nil, err
	}

	agent, err := NewGeTuiAgent(
		GTBaseUrl,
		conf.TestApp.AppId,
		conf.TestApp.AppKey,
		conf.TestApp.MasterSecret,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}

	err = mgr.RegisterAgent(UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}, agent)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	return &Server{
		Dao: dao,
		Mgr: mgr,
	}, nil
}

func (s *Server) GetToken(ctx context.Context, in *pb.GetTokenReq) (*pb.GetTokenResp, error) {
	out := &pb.GetTokenResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Dao.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("dao get token err %+v", err)
		return out, err
	}

	if auth == nil || in.OldToken == auth.Token || auth.ExpireAt < time.Now().Unix() {
		newAuth, err := s.Mgr.GetAuth(uniqueId)
		if err != nil {
			log.Errorf("mgr get token err %+v", err)
			return out, err
		}
		err = s.Dao.SetAuth(uniqueId, newAuth)
		if err != nil {
			log.Errorf("dao set token err %+v", err)
			return out, err
		}
		auth = newAuth
	}

	out.Token = auth.Token
	out.ExpireAt = auth.ExpireAt

	return out, nil
}

func (s *Server) SetToken(ctx context.Context, in *pb.SetTokenReq) (*pb.SetTokenResp, error) {
	out := &pb.SetTokenResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 || len(in.Token) == 0 || in.ExpireAt < time.Now().Unix() {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.Dao.SetAuth(uniqueId, &AuthToken{Token: in.Token, ExpireAt: in.ExpireAt})
	if err != nil {
		log.Errorf("dao set token err %+v", err)
		return out, err
	}

	return out, nil
}

func (s *Server) DelToken(ctx context.Context, in *pb.DelTokenReq) (*pb.DelTokenResp, error) {
	out := &pb.DelTokenResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 || len(in.Token) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.Mgr.DelAuth(uniqueId, in.Token)
	if err != nil {
		log.Errorf("mgr del token err %+v", err)
		return out, err
	}

	return out, nil
}

func (s *Server) Close() {
	s.Dao.Close()
}
