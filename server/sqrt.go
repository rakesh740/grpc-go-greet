package main

import (
	"context"
	"fmt"
	"math"
	pb "projects/Greet/proto"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (s *Server) Sqrt(ctx context.Context, in *pb.SqrtRequest) (*pb.SqrtResponse, error) {

	number := in.GetNum()

	if number < 0 {
		return nil, status.Errorf(codes.InvalidArgument, fmt.Sprintf("recieved a negetive number %f ", number))
	}
	return &pb.SqrtResponse{
		Answer: math.Sqrt(float64(number)),
	}, nil
}
