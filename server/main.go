package main

import (
	"net"

	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
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

	opts := []grpc.ServerOption{}

	tls := true // false if no security
	if tls {
		certFile := "ssl/server.crt"
		keyFile := "ssl/server.pem"
		creds, err := credentials.NewServerTLSFromFile(certFile, keyFile)
		if err != nil {
			logrus.Error(" error in loading certificates ")
		}

		opts = append(opts, grpc.Creds(creds))
	}

	s := grpc.NewServer(opts...)

	pb.RegisterCalculatorServer(s, &Server{})

	if err := s.Serve(lis); err != nil {
		logrus.Errorf(" error in serving in listener %v \n", err)
	}
}
