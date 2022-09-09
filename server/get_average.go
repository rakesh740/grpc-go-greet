package main

import (
	"io"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func (s *Server) GetAverage(stream pb.Calculator_GetAverageServer) error {

	logrus.Info(" GetAverage invoked ")
	var total, count int32
	for {
		msg, err := stream.Recv()

		if err == io.EOF {
			stream.SendAndClose(&pb.CalcResponse{
				Result: total / count,
			})
			logrus.Warn(" end of stream from client")
			return nil
		}

		if err != nil {
			logrus.Errorf(" error recieving from client side stream %v ", err)
			return nil
		}
		count++
		total += msg.GetNumber()
		logrus.Infof(" number %d, count %d", msg.GetNumber(), count)
	}
}
