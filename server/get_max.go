package main

import (
	"io"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func (s *Server) GetMaxNumber(stream pb.Calculator_GetMaxNumberServer) error {

	logrus.Info(" invoking GetMaxNumber ")

	var Max int32
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			logrus.Info(" ending streaming ")
			return nil
		}

		if msg.GetNumber() > Max {
			Max = msg.GetNumber()
			stream.Send(&pb.CalcResponse{
				Result: Max,
			})
		}
	}

}
