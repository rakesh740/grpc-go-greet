package main

import (
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
)

func (s *Server) GetAllPrimes(in *pb.NumberRequest, stream pb.Calculator_GetAllPrimesServer) error {

	logrus.Infof(" got request for GetAllPrimes with %d", in.GetNumber())
	var k int32 = 2
	number := in.GetNumber()

	for number > 1 {
		if number%k == 0 {

			stream.Send(&pb.NumberResponse{
				Prime: k,
			})
			time.Sleep(time.Second * 1)
			number = number / k
		} else {
			k = k + 1
		}
	}

	return nil
}
