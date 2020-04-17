package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"

	"github.com/infinity-oj/judger-sql/interval/baseinterface/test"

	"github.com/infinity-oj/judger-sql/interval/basestruct/report"

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

		log.Printf("Get judgement, token: %s, test point: %s", res.GetToken(), res.GetTestCase())
		token := res.GetToken()
		sampleFileName := fmt.Sprintf("%s.yaml", res.GetTestCase())
		sampleFile, err := getFile(c, token, "private", sampleFileName)
		if err != nil {
			log.Fatalf("could not get judgement file: %v", err)
		}

		sample, err := sample2.LoadFromBytes(sampleFile)
		if err != nil {
			log.Fatalf("could not parse judgement file: %v", err)
		}

		databaseDockerFile, err := getFile(c, token, "private", sample.Spec.DockerFile)
		if err != nil {
			log.Fatalf("could not find dockerfile : %v", err)
		}

		err = ioutil.WriteFile(sample.Spec.DockerFile, databaseDockerFile, 0666)
		if err != nil {
			log.Fatalf("could not write dockerfile : %v", err)
		}

		databaseBackup, err := getFile(c, token, "private", sample.Spec.Database)
		if err != nil {
			log.Fatalf("could not find database file : %v", err)
		}

		err = ioutil.WriteFile(sample.Spec.Database, databaseBackup, 0666)
		if err != nil {
			log.Fatalf("could not write database file : %v", err)
		}

		waitForPG, err := getFile(c, token, "private", "dockerfile/assignment3/wait-for-pg.sh")
		if err != nil {
			log.Fatalf("could not find database file : %v", err)
		}

		err = ioutil.WriteFile("dockerfile/assignment3/wait-for-pg.sh", waitForPG, 0666)
		if err != nil {
			log.Fatalf("could not write database file : %v", err)
		}

		userFileName := fmt.Sprintf("%s.sql", res.GetTestCase())
		userFile, err := getFile(c, token, "user", userFileName)
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
		t.Close()

		fmt.Println(r.SID)
		fmt.Println(r.Summary)
		fmt.Println(r.Grade)
		status, err := reportResult(c, token, uint64(r.Grade), r.Summary)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println(status)

		}
		time.Sleep(time.Second * 5)
	}
}

func getFile(client pb.JudgementsClient, token, spaceType, fileName string) ([]byte, error) {
	req := &pb.FetchJudgeFileRequest{
		Token:    token,
		Space:    spaceType,
		Filename: fileName,
	}

	res, err := client.FetchFile(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return res.GetFile(), nil
}

func reportResult(client pb.JudgementsClient, token string, grade uint64, msg string) (*pb.Status, error) {
	req := &pb.ReturnJudgementRequest{
		Token:  token,
		Status: 0,
		Score:  grade,
		Msg:    msg,
	}

	res, err := client.ReturnJudgement(context.TODO(), req)
	if err != nil {
		return nil, err
	}
	return &res.Status, err
}
