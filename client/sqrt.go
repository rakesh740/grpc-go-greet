package main

import (
	"context"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/status"
)

func doSqrt(c pb.CalculatorClient) {
	logrus.Info(" calling doSqrt")

	res, err := c.Sqrt(context.Background(), &pb.SqrtRequest{
		Num: -4,
	})

	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			logrus.Errorf(" error from server code  %s and message %s ", e.Code(), e.Message())
		} else {
			logrus.Error(" non grpc error ", err)
		}
	}

	logrus.Info(res.Answer)
}
