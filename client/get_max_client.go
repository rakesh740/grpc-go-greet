package main

import (
	"context"
	"io"
	pb "projects/Greet/proto"
	"time"

	"github.com/sirupsen/logrus"
)

func doGetMax(c pb.CalculatorClient) {

	logrus.Info(" calling do Get Max ")

	allnums := []*pb.NumberRequest{
		{Number: 1},
		{Number: 5},
		{Number: 3},
		{Number: 6},
		{Number: 2},
		{Number: 20},
	}

	stream, err := c.GetMaxNumber(context.Background())
	if err != nil {
		logrus.Error(" error while calling GetAverage ")
	}

	waitc := make(chan struct{})

	go func() {
		for _, v := range allnums {
			logrus.Infof(" sending number %d ", v.GetNumber())
			stream.Send(v)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func() {
		for {
			msg, err := stream.Recv()
			if err == io.EOF {
				logrus.Warn(" end of stream ")
				break
			}
			if err != nil {
				logrus.Errorf(" error in recieving max number from stream %v \n", err)
			}
			logrus.Printf("max number yet %d ", msg.GetResult())
		}
		close(waitc)
	}()

	<-waitc

}
