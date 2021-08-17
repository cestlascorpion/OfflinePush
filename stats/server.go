package stats

import (
	"context"
	"errors"
	"time"

	. "github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	Dao  *StatsDao
	Mgr  *AgentMgr
	Auth *AuthCache
}

func NewServer(conf *StatsConfig) (*Server, error) {
	dao, err := NewStatsDao(conf)
	if err != nil {
		log.Errorf("new stats dao err %+v", err)
		return nil, err
	}

	mgr, err := NewAgentMgr()
	if err != nil {
		log.Errorf("new stats mgr err %+v", err)
		return nil, err
	}

	agent, err := NewGeTuiStats(
		GTBaseUrl,
		conf.TestApp.AppId,
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

	auth, err := NewAuthCache()
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
		Dao:  dao,
		Mgr:  mgr,
		Auth: auth,
	}, nil
}

func (s *Server) GetTasks(ctx context.Context, in *proto.GetTasksReq) (*proto.GetTasksResp, error) {
	out := &proto.GetTasksResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetTasks(uniqueId, in.TaskList, auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetTasks(uniqueId, in.TaskList, auth.Token)
		}
		if err != nil {
			log.Errorf("get tasks err %+v", err)
			return out, err
		}
	}
	if len(resp) != len(in.TaskList) {
		log.Warnf("some task is missing")
	}
	for task, statics := range resp {
		out.TaskList = append(out.TaskList, task)
		out.StaticsList = append(out.StaticsList, statics)
	}
	return out, nil
}

func (s *Server) GetTaskGroup(ctx context.Context, in *proto.GetTaskGroupReq) (*proto.GetTaskGroupResp, error) {
	out := &proto.GetTaskGroupResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Group) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetTaskGroup(uniqueId, in.Group, auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetTaskGroup(uniqueId, in.Group, auth.Token)
		}
		if err != nil {
			log.Errorf("get task group err %+v", err)
			return out, err
		}
	}
	if len(resp) != 1 {
		log.Warnf("some group is missing")
	}
	for group, statics := range resp {
		out.Group = group
		out.Statics = statics
	}
	return out, nil
}

func (s *Server) GetPushCount(ctx context.Context, in *proto.GetPushCountReq) (*proto.GetPushCountResp, error) {
	out := &proto.GetPushCountResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") != time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetPushCount(uniqueId, auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetPushCount(uniqueId, auth.Token)
		}
		if err != nil {
			log.Errorf("get push count err %+v", err)
			return out, err
		}
	}
	for i := range resp {
		out.CountList = append(out.CountList, resp[i])
	}
	return out, nil
}

func (s *Server) GetPushDataByDay(ctx context.Context, in *proto.GetPushDataByDayReq) (*proto.GetPushDataByDayResp, error) {
	out := &proto.GetPushDataByDayResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") == time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetPushDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetPushDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
		}
		if err != nil {
			log.Errorf("get push data by day err %+v", err)
			return out, err
		}
	}
	if len(resp) != 1 {
		log.Warnf("something is missing")
	}
	for date, statics := range resp {
		out.Date = date
		out.Statics = statics
	}
	return out, nil
}

func (s *Server) GetUserDataByDay(ctx context.Context, in *proto.GetUserDataByDayReq) (*proto.GetUserDataByDayResp, error) {
	out := &proto.GetUserDataByDayResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") == time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetUserDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetUserDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
		}
		if err != nil {
			log.Errorf("get user data by day err %+v", err)
			return out, err
		}
	}
	if len(resp) != 1 {
		log.Warnf("something is missing")
	}
	for date, statics := range resp {
		out.Date = date
		for key, value := range statics {
			switch key {
			case "accumulative_num":
				out.Accumulative = value
			case "register_num":
				out.Register = value
			case "active_num":
				out.Active = value
			case "online_num":
				out.Online = value
			}
		}
	}
	return out, nil
}

func (s *Server) GetOnlineUserBy24H(ctx context.Context, in *proto.GetOnlineUserBy24HReq) (*proto.GetOnlineUserBy24HResp, error) {
	out := &proto.GetOnlineUserBy24HResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") != time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetOnlineUserBy24H(uniqueId, auth.Token)
	if err != nil {
		if err.Error() == InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.GetOnlineUserBy24H(uniqueId, auth.Token)
		}
		if err != nil {
			log.Errorf("get online data by day err %+v", err)
			return out, err
		}
	}
	for timestamp, statics := range resp {
		out.OnlineList = append(out.OnlineList, &proto.GetOnlineUserBy24HResp_OnlineInfo{
			UnixMillisecond: timestamp,
			Online:          statics,
		})
	}
	return out, nil
}

func (s *Server) Close() {
	s.Dao.Close()
}
