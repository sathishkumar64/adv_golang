package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/sathishkumar64/adv_golang/grpc_exmp"

	"google.golang.org/grpc"
)

func main() {
	con, err := grpc.Dial(":8080", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server is not reachable :%v", err)
	}
	client := grpc_exmp.NewGetMyBuildClient(con)
	err = getBuild(context.Background(), client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func getBuild(ctx context.Context, client grpc_exmp.GetMyBuildClient) error {

	buildRequest := grpc_exmp.BuildRequest{
		BuildId: "001",
	}

	buildRes, err := client.MyBuildInfo(ctx, &buildRequest)

	if err != nil {
		return fmt.Errorf("Could not fetch task: %v", err)
	}
	//	fmt.Printf("The output is:::: %v %v %v", buildRes.GetBuildId(), buildRes.GetBuildstatus(), buildRes.GetOsName())
	fmt.Printf("The output is:::: %v ", buildRes)

	return nil
}
