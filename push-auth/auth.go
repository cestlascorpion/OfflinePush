package main

import (
	"net"

	"github.com/cestlascorpion/push/auth"
	"github.com/cestlascorpion/push/core"
	"github.com/cestlascorpion/push/proto"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", core.AuthServerAddr)
	if err != nil {
		log.Fatalf("listen failed err %+v", err)
		return
	}
	conf := &core.AuthConfig{}
	err = configor.Load(conf, "auth.yml")
	if err != nil {
		log.Fatalf("config failed err %+v", err)
		return
	}
	svr, err := auth.NewServer(conf)
	if err != nil {
		log.Fatalf("new server failed err %+v", err)
		return
	}
	defer svr.Close()

	s := grpc.NewServer()
	proto.RegisterAuthServer(s, svr)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve failed err %+v", err)
		return
	}
}
