package main

import (
	"fmt"
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Server) GreetManyResponse(in *pb.GreetRequest, stream pb.Calculator_GreetManyResponseServer) error {

	logrus.Info(" got request for GreetManyResponse")

	for i := 0; i < 9; i++ {
		res := fmt.Sprintf("welcome to the world %s %d time/s \n", in.GetName(), i)
		stream.Send(&pb.GreetResponse{
			Response: res,
		})

		time.Sleep(time.Second * 1)
	}

	return nil
}
