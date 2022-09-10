package main

import (
	"context"
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func doGreetWithDeadline(c pb.CalculatorClient) {
	logrus.Info(" doGreetWithDeadline invoked")

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*2)
	defer cancel()

	req := pb.GreetRequest{
		Name: " Rakesh",
	}

	res, err := c.GreetWithDeadline(ctx, &req)
	if err != nil {
		e, ok := status.FromError(err)
		if ok {
			if e.Code() == codes.DeadlineExceeded {
				logrus.Error(" deadline exceeded ")
			} else {
				logrus.Error(" unexpected grpc error ")
			}
		} else {
			logrus.Error(" non grpc error ")
		}

		return
	}

	logrus.Info(res.Response, " got response in doGreetWithDeadline")
}
