package main

import (
	"context"
	"fmt"
	"log"
	"time"
)

func main() {

	ctx := context.Background()
	ctx, cancle := context.WithTimeout(ctx, time.Second)
	//efer cancle()
	cancle()

	// how to sending ctx values.
	mytestMethod(ctx, 5*time.Second, "context test")
}

func mytestMethod(ctx context.Context, d time.Duration, msg string) {
	//time.Sleep(d)
	select {
	case <-time.After(d):
		fmt.Println(msg)
	case <-ctx.Done():
		log.Print(ctx.Err())
	}
}
