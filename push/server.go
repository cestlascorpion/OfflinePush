package push

import (
	"context"
	"github.com/cestlascorpion/offlinepush/core"
	pb "github.com/cestlascorpion/offlinepush/proto"
)

type Server struct {
	*pb.UnimplementedPushServer
}

func NewServer(conf *core.PushConfig) (*Server, error) {
	return &Server{}, nil
}

func (s *Server) PushToSingle(context.Context, *pb.PushToSingleReq) (*pb.PushToSingleResp, error) {
	out := &pb.PushToSingleResp{}

	return out, nil
}

func (s *Server) CreateTask(context.Context, *pb.CreateTaskReq) (*pb.CreateTaskResp, error) {
	out := &pb.CreateTaskResp{}

	return out, nil
}

func (s *Server) PushToList(context.Context, *pb.PushToListRep) (*pb.PushToListResp, error) {
	out := &pb.PushToListResp{}

	return out, nil
}

func (s *Server) PushToApp(context.Context, *pb.PushToAppReq) (*pb.PushToAppResp, error) {
	out := &pb.PushToAppResp{}

	return out, nil
}

func (s *Server) StopTask(context.Context, *pb.StopTaskReq) (*pb.StopTaskResp, error) {
	out := &pb.StopTaskResp{}

	return out, nil
}

func (s *Server) CheckTask(context.Context, *pb.CheckTaskReq) (*pb.CheckTaskResp, error) {
	out := &pb.CheckTaskResp{}

	return out, nil
}

func (s *Server) RemoveTask(context.Context, *pb.RemoveTaskReq) (*pb.RemoveTaskResp, error) {
	out := &pb.RemoveTaskResp{}

	return out, nil
}

func (s *Server) ViewDetail(context.Context, *pb.ViewDetailReq) (*pb.ViewDetailResp, error) {
	out := &pb.ViewDetailResp{}

	return out, nil
}

func (s *Server) Close() {}
