package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"google.golang.org/grpc"

	"github.com/sathishkumar64/adv_golang/todo"
)

func main() {
	con, err := grpc.Dial(":8888", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Server is not reachable :%v", err)
	}
	client := todo.NewTasksClient(con)
	err = list(context.Background(), client)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func list(ctx context.Context, client todo.TasksClient) error {
	list, err := client.List(ctx, &todo.Void{})

	//log.Println(err)

	if err != nil {
		return fmt.Errorf("Could not fetch task: %v", err)
	}

	for _, t := range list.Tasks {

		if t.Done {
			fmt.Printf("Done")
		} else {
			fmt.Printf("No something went bad")
		}
		fmt.Printf("%s\n", t.Text)
	}
	return nil
}
