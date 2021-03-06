package push

import (
	"context"
	"errors"
	"net/http"
	"strconv"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/http2"
)

type Server struct {
	*proto.UnimplementedPushServer
	mgr  *AgentMgr
	auth *core.AuthCache
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	mgr, err := NewAgentMgr()
	if err != nil {
		log.Errorf("new stats mgr err %+v", err)
		return nil, err
	}

	g, err := NewGeTuiPush(
		core.GTBaseUrl,
		conf.GeTui.AppId,
		&http.Client{
			Transport: &http.Transport{},
		})
	if err != nil {
		log.Errorf("new getui agent err %+v", err)
		return nil, err
	}

	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.GeTui.AgentId, BundleId: conf.GeTui.BundleId}, g)
	if err != nil {
		log.Errorf("register getui agent err %+v", err)
		return nil, err
	}

	a, err := NewApnsPush(
		conf.Apns.Env,
		conf.Apns.Key,
		&http.Client{
			Transport: &http2.Transport{},
		})
	if err != nil {
		log.Errorf("new apns agent err %+v", err)
		return nil, err
	}

	err = mgr.RegisterAgent(core.UniqueId{PushAgent: conf.Apns.AgentId, BundleId: conf.Apns.BundleId}, a)
	if err != nil {
		log.Errorf("register apns agent err %+v", err)
		return nil, err
	}

	auth, err := core.NewAuthCache()
	if err != nil {
		log.Errorf("new auth cache err %+v", err)
		return nil, err
	}

	return &Server{
		mgr:  mgr,
		auth: auth,
	}, nil
}

func (s *Server) PushToSingle(ctx context.Context, in *proto.PushToSingleReq) (*proto.PushToSingleResp, error) {
	out := &proto.PushToSingleResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.MsgList) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	if len(in.MsgList) == 1 {
		err := s.pushSingle(ctx, uniqueId, in, out)
		if err != nil {
			log.Errorf("push single err %+v", err)
			return out, err
		}
		return out, nil
	}

	err := s.pushBatch(ctx, uniqueId, in, out)
	if err != nil {
		log.Errorf("push batch err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) CreateTask(ctx context.Context, in *proto.CreateTaskReq) (*proto.CreateTaskResp, error) {
	out := &proto.CreateTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.createTask(ctx, uniqueId, in, out)
	if err != nil {
		log.Errorf("create task err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) PushToList(ctx context.Context, in *proto.PushToListReq) (*proto.PushToListResp, error) {
	out := &proto.PushToListResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	err := s.pushList(ctx, uniqueId, in, out)
	if err != nil {
		log.Errorf("push list err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) PushToApp(ctx context.Context, in *proto.PushToAppReq) (*proto.PushToAppResp, error) {
	out := &proto.PushToAppResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		in.Msg == nil {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	if in.Msg.Audience != nil {
		err := s.pushByTag(ctx, uniqueId, in, out)
		if err != nil {
			log.Errorf("push by tag err %+v", err)
			return out, err
		}
		return out, nil
	}
	err := s.pushAll(ctx, uniqueId, in, out)
	if err != nil {
		log.Errorf("push all err %+v", err)
		return out, err
	}
	return out, nil
}

func (s *Server) StopTask(ctx context.Context, in *proto.StopTaskReq) (*proto.StopTaskResp, error) {
	out := &proto.StopTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	success, err := s.mgr.StopPush(ctx, uniqueId, in.TaskId, auth.Token)
	if err != nil {
		log.Errorf("stop push err %+v", err)
		return out, err
	}
	out.Success = success
	if !success {
		log.Warnf("stop push task failed")
	}
	return out, nil
}

func (s *Server) CheckTask(ctx context.Context, in *proto.CheckTaskReq) (*proto.CheckTaskResp, error) {
	out := &proto.CheckTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryScheduleTask(ctx, uniqueId, in.TaskId, auth.Token)
	if err != nil {
		log.Errorf("query shcedule task err %+v", err)
		return out, err
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

func (s *Server) RemoveTask(ctx context.Context, in *proto.RemoveTaskReq) (*proto.RemoveTaskResp, error) {
	out := &proto.RemoveTaskResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	success, err := s.mgr.DeleteScheduleTask(ctx, uniqueId, in.TaskId, auth.Token)
	if err != nil {
		log.Errorf("delete schedule task err %+v", err)
		return out, err
	}
	out.Success = success
	if !success {
		log.Warnf("delete shcedule task failed")
	}
	return out, nil
}

func (s *Server) ViewDetail(ctx context.Context, in *proto.ViewDetailReq) (*proto.ViewDetailResp, error) {
	out := &proto.ViewDetailResp{}

	if len(in.PushAgent) == 0 || len(in.BundleId) == 0 ||
		len(in.TaskId) == 0 || len(in.Cid) == 0 {
		log.Errorf("invalid parameter in %+v", in)
		return out, errors.New("invalid parameter")
	}

	uniqueId := core.UniqueId{PushAgent: in.PushAgent, BundleId: in.BundleId}
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return out, err
	}
	resp, err := s.mgr.QueryDetail(ctx, uniqueId, in.TaskId, in.Cid, auth.Token)
	if err != nil {
		log.Errorf("query task detail err %+v", err)
		return out, err
	}
	for i := range resp {
		out.DetailList = append(out.DetailList, &proto.ViewDetailResp_Detail{
			Time:  resp[i][0],
			Event: resp[i][1],
		})
	}
	return out, nil
}

func (s *Server) Close() {
	s.auth.Close()
	s.mgr.Close()
}

func (s *Server) pushSingle(ctx context.Context, uniqueId core.UniqueId, in *proto.PushToSingleReq, out *proto.PushToSingleResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
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

	resp, err := s.mgr.PushSingle(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("push single err %+v", err)
		return err
	}

	if len(resp) != 1 {
		log.Errorf("push single unexpected resp %+v", resp)
		return errors.New("invalid push single resp")
	}

	out.ReceiptList = make([]*proto.Receipt, 0)
	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*proto.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &proto.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.ReceiptList = append(out.ReceiptList, &proto.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		})
	}
	return nil
}

func (s *Server) pushBatch(ctx context.Context, uniqueId core.UniqueId, in *proto.PushToSingleReq, out *proto.PushToSingleResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
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

	resp, err := s.mgr.PushBatch(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("push batch err %+v", err)
		return err
	}

	out.ReceiptList = make([]*proto.Receipt, 0)
	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*proto.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &proto.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.ReceiptList = append(out.ReceiptList, &proto.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		})
	}
	return nil
}

func (s *Server) createTask(ctx context.Context, uniqueId core.UniqueId, in *proto.CreateTaskReq, out *proto.CreateTaskResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
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

	taskId, err := s.mgr.CreateMsg(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("create msg err %+v", err)
		return err
	}
	out.TaskId = taskId
	return nil
}

func (s *Server) pushList(ctx context.Context, uniqueId core.UniqueId, in *proto.PushToListReq, out *proto.PushToListResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
	if err != nil {
		log.Errorf("get auth err %+v", err)
		return err
	}
	req := &ListReq{
		Audience: in.Msg.Audience,
		IsAsync:  in.IsAsync,
		TaskId:   in.Msg.TaskId,
	}

	resp, err := s.mgr.PushList(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("push list err %+v", err)
		return err
	}

	if len(resp) != 1 {
		log.Errorf("push list unexpected resp %+v", resp)
		return errors.New("invalid push list resp")
	}

	for taskId, detail := range resp {
		if len(detail) == 0 {
			continue
		}
		detailList := make([]*proto.Receipt_Detail, 0)
		for cid, status := range detail {
			detailList = append(detailList, &proto.Receipt_Detail{
				Cid:    cid,
				Status: status,
			})
		}
		out.Receipt = &proto.Receipt{
			TaskId:     taskId,
			DetailList: detailList,
		}
	}
	return nil
}

func (s *Server) pushAll(ctx context.Context, uniqueId core.UniqueId, in *proto.PushToAppReq, out *proto.PushToAppResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
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

	taskId, err := s.mgr.PushAll(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("push all err %+v", err)
		return err
	}
	out.TaskId = taskId
	return nil
}

func (s *Server) pushByTag(ctx context.Context, uniqueId core.UniqueId, in *proto.PushToAppReq, out *proto.PushToAppResp) error {
	auth, err := s.auth.GetAuth(uniqueId)
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

	taskId, err := s.mgr.PushByTag(ctx, uniqueId, req, auth.Token)
	if err != nil {
		log.Errorf("push by tag err %+v", err)
		return err
	}
	out.TaskId = taskId
	return nil
}
