package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "projects/Greet/proto"
)

var addr string = "0.0.0.0:8192"

func main() {

	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf(" error in dial connection %s", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	// doGreetManyTimes(c)

	// doSum(c)
	//doGetPrimes(c)
	doGetAverage(c)
}
