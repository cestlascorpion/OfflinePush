package push

import (
	"context"
	"errors"
	"strconv"
	"time"

	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
)

const (
	MaxBatchTarget = 200
	MaxListTarget  = 1000
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
		len(in.MsgList) == 0 || len(in.MsgList) > MaxBatchTarget {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	if len(in.MsgList) == 1 {
		err := s.pushSingle(uniqueId, in, out)
		if err != nil {
			log.Errorf("push single err %+v", err)
			return out, err
		}
		return out, nil
	}

	err := s.pushBatch(uniqueId, in, out)
	if err != nil {
		log.Errorf("push batch err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) CreateTask(ctx context.Context, in *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	out := &pb.CreateTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.createTask(uniqueId, in, out)
	if err != nil {
		log.Errorf("create task err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) PushToList(ctx context.Context, in *pb.PushToListReq) (*pb.PushToListResp, error) {
	out := &pb.PushToListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil || len(in.Msg.Audience.Cid) > MaxListTarget || len(in.Msg.Audience.Alias) > MaxListTarget {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.pushList(uniqueId, in, out)
	if err != nil {
		log.Errorf("push list err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) PushToApp(ctx context.Context, in *pb.PushToAppReq) (*pb.PushToAppResp, error) {
	out := &pb.PushToAppResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	if in.Msg.Audience != nil {
		err := s.pushByTag(uniqueId, in, out)
		if err != nil {
			log.Errorf("push by tag err %+v", err)
			return out, err
		}
		return out, nil
	}
	err := s.pushAll(uniqueId, in, out)
	if err != nil {
		log.Errorf("push all err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) StopTask(ctx context.Context, in *pb.StopTaskReq) (*pb.StopTaskResp, error) {
	out := &pb.StopTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	success, err := s.Mgr.StopPush(uniqueId, in.TaskId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			success, err = s.Mgr.StopPush(uniqueId, in.TaskId, auth.Token)
		}
		if err != nil {
			log.Errorf("stop push err %+v", err)
			return out, err
		}
	}
	out.Success = success
	if !success {
		log.Warnf("stop push task failed")
	}
	return out, nil
}

func (s *Server) CheckTask(ctx context.Context, in *pb.CheckTaskReq) (*pb.CheckTaskResp, error) {
	out := &pb.CheckTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.Mgr.QueryScheduleTask(uniqueId, in.TaskId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			resp, err = s.Mgr.QueryScheduleTask(uniqueId, in.TaskId, auth.Token)
		}
		if err != nil {
			log.Errorf("query shcedule task err %+v", err)
			return out, err
		}
	}
	for k, v := range resp {
		switch k {
		case "create_time":
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Errorf("strconv parse %s -> int err %+v", v, err)
				continue
			}
			out.CreateTime = ts
		case "status":
			out.Status = v
		case "transmission_content":
			out.TransmissionContent = v
		case "push_time":
			ts, err := strconv.ParseInt(v, 10, 64)
			if err != nil {
				log.Errorf("strconv parse %s -> int err %+v", v, err)
				continue
			}
			out.PushTime = ts
		}
	}
	return out, nil
}

func (s *Server) RemoveTask(ctx context.Context, in *pb.RemoveTaskReq) (*pb.RemoveTaskResp, error) {
	out := &pb.RemoveTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	success, err := s.Mgr.DeleteScheduleTask(uniqueId, in.TaskId, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return out, err
			}
			success, err = s.Mgr.DeleteScheduleTask(uniqueId, in.TaskId, auth.Token)
		}
		if err != nil {
			log.Errorf("delete schedule task err %+v", err)
			return out, err
		}
	}
	out.Success = success
	if !success {
		log.Warnf("delete shcedule task failed")
	}
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
			log.Errorf("query task detail err %+v", err)
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

func (s *Server) pushSingle(uniqueId core.UniqueId, in *pb.PushToSingleReq, out *pb.PushToSingleResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &SingleReq{
		RequestId:   in.MsgList[0].RequestId,
		Audience:    in.MsgList[0].Audience,
		Settings:    in.MsgList[0].Settings,
		PushMessage: in.MsgList[0].PushMessage,
		PushChannel: in.MsgList[0].PushChannel,
	}

	resp, err := s.Mgr.PushSingle(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			resp, err = s.Mgr.PushSingle(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("push single err %+v", err)
			return err
		}
	}

	if len(resp) != 1 {
		log.Errorf("push single unexpected resp %+v", resp)
		return errors.New("invalid push single resp")
	}

	out.ReceiptList = make([]*pb.Receipt, 0)
	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*pb.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &pb.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.ReceiptList = append(out.ReceiptList, &pb.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		})
	}
	return nil
}

func (s *Server) pushBatch(uniqueId core.UniqueId, in *pb.PushToSingleReq, out *pb.PushToSingleResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &BatchReq{
		IsAsync: in.IsAsync,
		MsgList: make([]*SingleReq, 0),
	}
	for i := range in.MsgList {
		req.MsgList = append(req.MsgList, &SingleReq{
			RequestId:   in.MsgList[i].RequestId,
			Audience:    in.MsgList[i].Audience,
			Settings:    in.MsgList[i].Settings,
			PushMessage: in.MsgList[i].PushMessage,
			PushChannel: in.MsgList[i].PushChannel,
		})
	}

	resp, err := s.Mgr.PushBatch(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			resp, err = s.Mgr.PushBatch(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("push batch err %+v", err)
			return err
		}
	}

	out.ReceiptList = make([]*pb.Receipt, 0)
	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*pb.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &pb.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.ReceiptList = append(out.ReceiptList, &pb.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		})
	}
	return nil
}

func (s *Server) createTask(uniqueId core.UniqueId, in *pb.CreateTaskReq, out *pb.CreateTaskResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &CreateReq{
		RequestId:   in.Msg.RequestId,
		GroupName:   in.Msg.GroupName,
		Settings:    in.Msg.Settings,
		PushMessage: in.Msg.PushMessage,
		PushChannel: in.Msg.PushChannel,
	}

	taskId, err := s.Mgr.CreateMsg(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			taskId, err = s.Mgr.CreateMsg(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("create msg err %+v", err)
			return err
		}
	}
	out.TaskId = taskId
	return nil
}

func (s *Server) pushList(uniqueId core.UniqueId, in *pb.PushToListReq, out *pb.PushToListResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &ListReq{
		Audience: in.Msg.Audience,
		IsAsync:  in.IsAsync,
		TaskId:   in.Msg.TaskId,
	}

	resp, err := s.Mgr.PushList(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			resp, err = s.Mgr.PushList(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("push list err %+v", err)
			return err
		}
	}

	if len(resp) != 1 {
		log.Errorf("push list unexpected resp %+v", resp)
		return errors.New("invalid push list resp")
	}

	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*pb.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &pb.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.Receipt = &pb.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		}
	}
	return nil
}

func (s *Server) pushAll(uniqueId core.UniqueId, in *pb.PushToAppReq, out *pb.PushToAppResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &AllReq{
		RequestId:   in.Msg.RequestId,
		GroupName:   in.Msg.GroupName,
		Audience:    "all",
		Settings:    in.Msg.Settings,
		PushMessage: in.Msg.PushMessage,
		PushChannel: in.Msg.PushChannel,
	}

	taskId, err := s.Mgr.PushAll(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			taskId, err = s.Mgr.PushAll(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("push all err %+v", err)
			return err
		}
	}
	out.TaskId = taskId
	return nil
}

func (s *Server) pushByTag(uniqueId core.UniqueId, in *pb.PushToAppReq, out *pb.PushToAppResp) error {
	auth, err := s.Auth.GetAuth(uniqueId, "")
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &ByTagReq{
		RequestId:   in.Msg.RequestId,
		GroupName:   in.Msg.GroupName,
		Settings:    in.Msg.Settings,
		Audience:    in.Msg.Audience,
		PushMessage: in.Msg.PushMessage,
		PushChannel: in.Msg.PushChannel,
	}

	taskId, err := s.Mgr.PushByTag(uniqueId, req, auth.Token)
	if err != nil {
		if err.Error() == core.InvalidTokenErr {
			auth, err = s.Auth.GetAuth(uniqueId, auth.Token)
			if err != nil {
				log.Errorf("get authx2 err %+v", err)
				return err
			}
			taskId, err = s.Mgr.PushByTag(uniqueId, req, auth.Token)
		}
		if err != nil {
			log.Errorf("push by tag err %+v", err)
			return err
		}
	}
	out.TaskId = taskId
	return nil
}
