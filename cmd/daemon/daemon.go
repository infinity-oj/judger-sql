package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	pb "github.com/infinity-oj/judger-sql/api/protobuf-spec"
	"github.com/infinity-oj/judger-sql/interval/consul"
	"google.golang.org/grpc"
)

const (
	target      = "consul://192.168.31.211:8500/Judgements"
	defaultName = "world"
)

func main() {
	consul.Init()
	// Set up a connection to the server.
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)

	conn, err := grpc.DialContext(ctx, target, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewJudgementsClient(conn)

	// Contact the server and print out its response.
	name := defaultName
	if len(os.Args) > 1 {
		name = os.Args[1]
	}
	fmt.Print(name)
	for {
		ctx, _ := context.WithTimeout(context.Background(), time.Second)

		req := &pb.FetchJudgementRequest{
			TimeLimit:   0,
			MemoryLimit: 0,
		}

		res, err := c.FetchJudgement(ctx, req)
		if err != nil {
			log.Fatalf("could not get judgement: %v", err)
		}

		log.Printf("Get judgement, test point: %s", res.GetTestCase())

		time.Sleep(time.Duration(2) * time.Second)
	}
}
