package main

import (
	"context"
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) GreetWithDeadline(ctx context.Context, in *pb.GreetRequest) (*pb.GreetResponse, error) {

	logrus.Info(" GreetWithDeadline called ")

	for i := 0; i < 3; i++ {
		if ctx.Err() == context.DeadlineExceeded {
			logrus.Error(" client cancelled the request!")
			return nil, status.Error(codes.Canceled, "client cancelled the request! ")
		}

		time.Sleep(1 * time.Second)
	}

	return &pb.GreetResponse{
		Response: "Hello " + in.Name,
	}, nil
}
