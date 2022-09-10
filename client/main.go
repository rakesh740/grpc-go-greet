package main

import (
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"

	pb "projects/Greet/proto"
)

var addr string = "localhost:8192"

func main() {
	tls := true
	opts := []grpc.DialOption{}

	if tls {
		certFile := "ssl/ca.crt"
		creds, err := credentials.NewClientTLSFromFile(certFile, "")
		if err != nil {
			logrus.Error(" error while loading certificate in server ")
		}

		opts = append(opts, grpc.WithTransportCredentials(creds))
	}

	conn, err := grpc.Dial(addr, opts...)
	if err != nil {
		logrus.Errorf(" error in dial connection %s", err)
	}
	defer conn.Close()

	c := pb.NewCalculatorClient(conn)

	doGreetManyTimes(c)

	// doSum(c)
	//doGetPrimes(c)
	//doGetAverage(c)
	//doGetMax(c)

	//doSqrt(c)
	//doGreetWithDeadline(c)
}
