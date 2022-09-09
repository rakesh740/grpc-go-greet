package main

import (
	"context"
	pb "projects/Greet/proto"

	"github.com/sirupsen/logrus"
)

func doSum(c pb.CalculatorClient) {

	logrus.Info(" sum was called from client ")

	res, err := c.Sum(context.Background(), &pb.Operand{FirstNum: 5, SecondNum: 1})
	if err != nil {
		logrus.Errorf(" could not call sum %v\n ", err)
	}

	logrus.Infof(" result %v ", res)

}
