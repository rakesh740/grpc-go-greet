package main

import (
	"context"
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
)

func doGetAverage(c pb.CalculatorClient) {
	logrus.Info("calling doGetAverage")

	allnums := []*pb.NumberRequest{
		{Number: 1},
		{Number: 9},
		{Number: 2},
		{Number: 8},
		{Number: 3},
		{Number: 7},
	}

	stream, err := c.GetAverage(context.Background())
	if err != nil {
		logrus.Error(" error while calling GetAverage ")
	}

	for _, v := range allnums {
		stream.Send(v)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		logrus.Error(" error while recieving response from GetAverage ")
	}

	logrus.Info(" result is ", res.GetResult())
}
