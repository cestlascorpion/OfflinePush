package main

import (
	"net"

	"github.com/cestlascorpion/offlinepush/core"
	"github.com/cestlascorpion/offlinepush/proto"
	"github.com/cestlascorpion/offlinepush/stats"
	"github.com/jinzhu/configor"
	log "github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	lis, err := net.Listen("tcp", core.StatsServerAddr)
	if err != nil {
		log.Fatalf("listen failed err %+v", err)
		return
	}
	conf := &core.PushConfig{}
	err = configor.Load(conf, "conf.yml")
	if err != nil {
		log.Fatalf("config failed err %+v", err)
		return
	}
	svr, err := stats.NewServer(conf)
	if err != nil {
		log.Fatalf("new server failed err %+v", err)
		return
	}
	defer svr.Close()

	s := grpc.NewServer()
	proto.RegisterStatsServer(s, svr)
	reflection.Register(s)

	err = s.Serve(lis)
	if err != nil {
		log.Fatalf("serve failed err %+v", err)
		return
	}
}
