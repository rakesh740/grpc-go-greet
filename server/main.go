package main

import (
	"net"

	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

type Server struct {
	pb.CalculatorServer
}

var addr string = "0.0.0.0:8192"

func main() {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		logrus.Errorf(" error in creating listener %v \n", err)
	}
	logrus.Infof("listening on address: %s\n", addr)

	s := grpc.NewServer()
	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		logrus.Errorf(" error in serving in listener %v \n", err)
	}
}
