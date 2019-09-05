package main

import (
	"context"
	"encoding/binary"
	"fmt"
	"log"
	"net"

	"github.com/sathishkumar64/adv_golang/todo"
	"google.golang.org/grpc"
)

type taskServer struct {
}

type length int64

const (
	sizeOfLength = 8
	dbPath       = "mydb.pb"
)

var endianness = binary.LittleEndian

func (t taskServer) List(ctx context.Context, void *todo.Void) (*todo.TaskList, error) {
	log.Println("Getting inside.................")
	return nil, fmt.Errorf("Not implemented")
}

func main() {
	srv := grpc.NewServer()

	var tasks taskServer
	todo.RegisterTasksServer(srv, tasks)
	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}
	log.Fatal(srv.Serve(lis))
}
