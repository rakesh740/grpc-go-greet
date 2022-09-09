package main

import (
	"context"
	"io"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func doGreetManyTimes(c pb.CalculatorClient) {
	logrus.Info("calling doGreetManyTimes")

	req := &pb.GreetRequest{
		Name: "Rakesh",
	}

	stream, err := c.GreetManyResponse(context.Background(), req)
	if err != nil {
		logrus.Error(" error in calling Greet many request %v", err)
	}

	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			logrus.Warn(" end of stream ")
			break
		}
		if err != nil {
			logrus.Errorf(" error in recieving greet response many times from stream %v \n", err)
		}
		logrus.Printf(msg.GetResponse())
	}
}
