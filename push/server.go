package push

import (
	"context"
	"errors"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*pb.UnimplementedPushServer
	Mgr  *AgentMgr
	Auth *core.AuthCache
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	mgr, err := NewAgentMgr()
	if err != nil {
		log.Errorf("new stats mgr err %+v", err)
		return nil, err
	}

	agent, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.TestApp.AppId,
		time.Duration(conf.TestApp.TimeoutSec)*time.Second)
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}

	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.TestApp.PushAgent, BundleId: conf.TestApp.BundleId}, agent)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	auth, err := core.NewAuthCache()
	if err != nil {
		log.Errorf("new auth cache err %+v", err)
		return nil, err
	}
	err = auth.Start(context.Background())
	if err != nil {
		log.Errorf("start auth cache err %+v", err)
		return nil, err
	}

	return &Server{
		Mgr:  mgr,
		Auth: auth,
	}, nil
}

func (s *Server) PushToSingle(ctx context.Context, in *pb.PushToSingleReq) (*pb.PushToSingleResp, error) {
	out := &pb.PushToSingleResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.MsgList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil // todo 分片
}

func (s *Server) CreateTask(ctx context.Context, in *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	out := &pb.CreateTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil
}

func (s *Server) PushToList(ctx context.Context, in *pb.PushToListRep) (*pb.PushToListResp, error) {
	out := &pb.PushToListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil // todo 分片
}

func (s *Server) PushToApp(ctx context.Context, in *pb.PushToAppReq) (*pb.PushToAppResp, error) {
	out := &pb.PushToAppResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil // todo 分片
}

func (s *Server) StopTask(ctx context.Context, in *pb.StopTaskReq) (*pb.StopTaskResp, error) {
	out := &pb.StopTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil
}

func (s *Server) CheckTask(ctx context.Context, in *pb.CheckTaskReq) (*pb.CheckTaskResp, error) {
	out := &pb.CheckTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil
}

func (s *Server) RemoveTask(ctx context.Context, in *pb.RemoveTaskReq) (*pb.RemoveTaskResp, error) {
	out := &pb.RemoveTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	// todo

	return out, nil
}

func (s *Server) ViewDetail(ctx context.Context, in *pb.ViewDetailReq) (*pb.ViewDetailResp, error) {
	out := &pb.ViewDetailResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 || len(in.Cid) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.QueryDetail(uniqueId, in.TaskId, in.Cid, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.QueryDetail(uniqueId, in.TaskId, in.Cid, auth.Token)
		}
		if err != nil {
			log.Errorf("get tasks err %+v", err)
			return out, err
		}
	}
	for i := range resp {
		out.DetailList = append(out.DetailList, &pb.ViewDetailResp_Detail{
			Time:  resp[i][0],
			Event: resp[i][1],
		})
	}
	return out, nil
}

func (s *Server) Close() {
	// do nothing
}
