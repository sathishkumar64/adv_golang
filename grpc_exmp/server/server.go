package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"google.golang.org/grpc/grpclog"

	"github.com/sathishkumar64/adv_golang/grpc_exmp"
	"google.golang.org/grpc"
)

// GetMyBuildServer Get My build Server
type GetMyBuildServer struct {
}

// MyBuildInfo Get my Build info
func (g GetMyBuildServer) MyBuildInfo(cxt context.Context, buildReq *grpc_exmp.BuildRequest) (*grpc_exmp.BuildResponse, error) {

	fmt.Printf("The input is %v", buildReq.GetBuildId())
	buildRes := grpc_exmp.BuildResponse{
		BuildId:     buildReq.GetBuildId(),
		OsName:      "Linux",
		Buildstatus: true,
	}
	return &buildRes, nil
}

var glog grpclog.LoggerV2

func init() {
	glog = grpclog.NewLoggerV2(os.Stdout, os.Stdout, os.Stdout)
}

func main() {

	srv := grpc.NewServer()
	var mybuild GetMyBuildServer
	grpc_exmp.RegisterGetMyBuildServer(srv, mybuild)
	lis, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Could not listen to :8080 :%v", err)
	}
	log.Fatal(srv.Serve(lis))

}
