package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/infinity-oj/judger-sql/interval/basestruct/report"

	"github.com/infinity-oj/judger-sql/interval/baseinterface/test"

	sample2 "github.com/infinity-oj/judger-sql/interval/service/sample"

	pb "github.com/infinity-oj/judger-sql/api/protobuf-spec"
	"github.com/infinity-oj/judger-sql/interval/consul"
	"google.golang.org/grpc"
)

const (
	target      = "consul://192.168.3.10:8500/Judgements"
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

		if res == nil || res.GetTestCase() == "" {
			log.Printf("nothing...")
			time.Sleep(time.Second * 5)
			continue
		}

		log.Printf("Get judgement, test point: %s", res.GetTestCase())

		sampleFileName := fmt.Sprintf("%s.yaml", res.GetTestCase())
		sampleFile, err := getFile(c, res.GetPrivateSpace(), sampleFileName)
		if err != nil {
			log.Fatalf("could not get judgement file: %v", err)
		}

		sample, err := sample2.LoadFromBytes(sampleFile)
		if err != nil {
			log.Fatalf("could not parse judgement file: %v", err)
		}

		userFileName := fmt.Sprintf("%s.sql", res.GetTestCase())
		userFile, err := getFile(c, res.GetUserSpace(), userFileName)
		if err != nil {
			log.Fatalf("could not get user file: %v", err)
		}

		t := test.SelectTest(*sample)
		err = t.Init(*sample, string(userFile))
		if err != nil {
			log.Fatalf("could not init test: %v", err)
		}

		reportChan := make(chan report.Report)
		go t.Run(reportChan)
		r := <-reportChan

		fmt.Println(r.SID)
		fmt.Println(r.Summary)
		fmt.Println(r.Grade)
	}
}

func getFile(client pb.JudgementsClient, fileSpace, fileName string) ([]byte, error) {
	req := &pb.FetchJudgeFileRequest{
		FileSpace: fileSpace,
		Filename:  fileName,
	}

	res, err := client.FetchFile(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return res.GetFile(), nil
}
