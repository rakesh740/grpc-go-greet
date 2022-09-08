package main

import (
	"context"

	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func (s *Server) Sum(ctx context.Context, in *pb.Operand) (*pb.Response, error) {

	logrus.Info(" performing addition of two numbers ")

	result := in.FirstNum + in.SecondNum

	return &pb.Response{
		Answer: result,
	}, nil
}
