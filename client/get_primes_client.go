package main

import (
	"context"
	"io"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func doGetPrimes(c pb.CalculatorClient) {
	logrus.Info("calling doGetPrimes")

	req := &pb.NumberRequest{
		Number: 120,
	}

	stream, err := c.GetAllPrimes(context.Background(), req)
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
			logrus.Errorf(" error in recieving prime numbers from stream %v \n", err)
		}
		logrus.Printf(" %d ", msg.GetPrime())
	}
}
