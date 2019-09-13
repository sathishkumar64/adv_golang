package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"
	"os/signal"

	"github.com/sathishkumar64/adv_golang/studentservice/model"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

//StduentServiceServer type is expose services.
type StduentServiceServer struct {
}

//CreateStu type is expose services.
func (s *StduentServiceServer) CreateStu(ctx context.Context, req *model.CreateStuReq) (*model.CreateStuRes, error) {
	std := req.GetStudent()

	data := Student{
		// ID:    Empty, so it gets omitted and MongoDB generates a unique Object ID upon insertion.
		studentID:   std.GetStudentId(),
		studentName: std.GetSchoolname(),
		className:   std.GetClassName(),
		schoolname:  std.GetSchoolname(),
	}
	_, err := studentdb.InsertOne(mongoCtx, data)
	// check for potential errors
	if err != nil {
		// return internal gRPC error to be handled later
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Internal error: %v", err),
		)
	}
	return &model.CreateStuRes{Student: std}, nil
}

//ReadStuBySchool type is expose services.
func (s *StduentServiceServer) ReadStuBySchool(ctx context.Context, req *model.ReadStuReq) (*model.ReadStuRes, error) {
	return nil, fmt.Errorf("Not implemented")
}

//ListStus type is expose services.
func (s *StduentServiceServer) ListStus(req *model.ListStuReq, stream model.StduentService_ListStusServer) error {
	return fmt.Errorf("Not implemented")
}

//Student type is expose services.
type Student struct {
	studentID   string `bson:"student_id,omitempty"`
	studentName string `bson:"student_name,omitempty"`
	className   string `bson:"class_name,omitempty"`
	schoolname  string `bson:"schoolname,omitempty"`
}

var db *mongo.Client
var studentdb *mongo.Collection
var mongoCtx context.Context

func main() {

	s := grpc.NewServer()
	srv := &StduentServiceServer{}
	model.RegisterStduentServiceServer(s, srv)

	lis, err := net.Listen("tcp", ":8888")
	if err != nil {
		log.Fatalf("Could not listen to :8888 :%v", err)
	}

	fmt.Println("Connecting to MongoDB...")
	mongoCtx = context.Background()
	db, err = mongo.Connect(mongoCtx, options.Client().ApplyURI("mongodb+srv://admin:admin@cluster0-hsyh0.gcp.mongodb.net"))
	if err != nil {
		log.Fatal(err)
	}
	err = db.Ping(mongoCtx, nil)
	if err != nil {
		log.Fatalf("Could not connect to MongoDB: %v\n", err)
	} else {
		fmt.Println("Connected to Mongodb")
	}

	studentdb = db.Database("studentservice").Collection("student")

	go func() {
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	fmt.Println("Server succesfully started on port :8888")
	c := make(chan os.Signal)

	// Relay os.Interrupt to our channel (os.Interrupt = CTRL+C)
	// Ignore other incoming signals
	signal.Notify(c, os.Interrupt)

	// Block main routine until a signal is received
	// As long as user doesn't press CTRL+C a message is not passed and our main routine keeps running
	<-c

	// After receiving CTRL+C Properly stop the server
	fmt.Println("\nStopping the server...")
	s.Stop()
	lis.Close()
	fmt.Println("Closing MongoDB connection")
	db.Disconnect(mongoCtx)
	fmt.Println("Done.")
}
