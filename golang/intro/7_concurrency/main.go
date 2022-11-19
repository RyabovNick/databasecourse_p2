package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/RyabovNick/databasecourse_2/golang/intro/7_concurrency/scan/chansync"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		shutdown := make(chan os.Signal, 1)
		signal.Notify(shutdown, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

		<-shutdown
		cancel()

		time.Sleep(10 * time.Second) // TODO: chansync

		os.Exit(0)
	}()

	ps := chansync.New("127.0.0.1", 100, 10)
	port := ps.Scan(ctx)
	fmt.Println("open ports:", port)
}
