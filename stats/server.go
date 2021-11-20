package stats

import (
	"context"
	"errors"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

type Server struct {
	*pb.UnimplementedStatsServer
	Dao  *core.StatsDao
	Mgr  *AgentMgr
	Auth *core.AuthCache
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	dao, err := core.NewStatsDao(conf)
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
		Dao:  dao,
		Mgr:  mgr,
		Auth: auth,
	}, nil
}

func (s *Server) GetTasks(ctx context.Context, in *pb.GetTasksReq) (*pb.GetTasksResp, error) {
	out := &pb.GetTasksResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetTasks(uniqueId, in.TaskList, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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

	if s.Dao.SetStats(uniqueId, "GetTasks", time.Now(), resp) != nil {
		log.Errorf("save tasks err %+v", err)
	}

	return out, nil
}

func (s *Server) GetTaskGroup(ctx context.Context, in *pb.GetTaskGroupReq) (*pb.GetTaskGroupResp, error) {
	out := &pb.GetTaskGroupResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.Group) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetTaskGroup(uniqueId, in.Group, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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

	if s.Dao.SetStats(uniqueId, "GetTaskGroup", time.Now(), resp) != nil {
		log.Errorf("save task group err %+v", err)
	}

	return out, nil
}

func (s *Server) GetPushCount(ctx context.Context, in *pb.GetPushCountReq) (*pb.GetPushCountResp, error) {
	out := &pb.GetPushCountResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetPushCount(uniqueId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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

	if s.Dao.SetStats(uniqueId, "GetPushCount", time.Now(), resp) != nil {
		log.Errorf("save push count err %+v", err)
	}

	return out, nil
}

func (s *Server) GetPushDataByDay(ctx context.Context, in *pb.GetPushDataByDayReq) (*pb.GetPushDataByDayResp, error) {
	out := &pb.GetPushDataByDayResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") == time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetPushDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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

	if s.Dao.SetStats(uniqueId, "GetPushDataByDay", time.Unix(in.UnixSecond, 0), resp) != nil {
		log.Errorf("save push data by day err %+v", err)
	}

	return out, nil
}

func (s *Server) GetUserDataByDay(ctx context.Context, in *pb.GetUserDataByDayReq) (*pb.GetUserDataByDayResp, error) {
	out := &pb.GetUserDataByDayResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		time.Unix(in.UnixSecond, 0).Format("2006-01-02") == time.Now().Format("2006-01-02") {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetUserDataByDay(uniqueId, time.Unix(in.UnixSecond, 0), auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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

	if s.Dao.SetStats(uniqueId, "GetUserDataByDay", time.Unix(in.UnixSecond, 0), resp) != nil {
		log.Errorf("save user data by day err %+v", err)
	}

	return out, nil
}

func (s *Server) GetOnlineUserBy24H(ctx context.Context, in *pb.GetOnlineUserBy24HReq) (*pb.GetOnlineUserBy24HResp, error) {
	out := &pb.GetOnlineUserBy24HResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.GetOnlineUserBy24H(uniqueId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
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
		out.OnlineList = append(out.OnlineList, &pb.GetOnlineUserBy24HResp_OnlineInfo{
			UnixMillisecond: timestamp,
			Online:          statics,
		})
	}

	if s.Dao.SetStats(uniqueId, "GetOnlineUserBy24H", time.Now(), resp) != nil {
		log.Errorf("save online user by 24h err %+v", err)
	}

	return out, nil
}

func (s *Server) Close() {
	s.Dao.Close()
}
