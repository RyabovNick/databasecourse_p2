package chansync

import (
	"context"
	"errors"
	"fmt"
	"net"
	"sync"
)

// TODO: подготовить Scan к отмене ctx, т.к в данный момент
// только return в worker, но wg.Wait() будет всегда ждать, пока
// wg == 0, но из-за отсутствия worker это никогда не случится, т.е.
// приложение будет вечно висеть

type PortScanner struct {
	// host to scan
	host string
	// port to scan to
	portTo int

	jobs   chan int
	result chan int
}

func New(host string, portTo int, worker int) *PortScanner {
	return &PortScanner{
		host:   host,
		portTo: portTo,
		jobs:   make(chan int, worker),
		result: make(chan int),
	}
}

func (ps *PortScanner) Scan(ctx context.Context) []int {
	// wg for worker jobs
	var wg sync.WaitGroup

	// start N workers
	for i := 0; i < cap(ps.jobs); i++ {
		go ps.worker(ctx, &wg)
	}

	// send ports to jobs
	for i := 1; i < ps.portTo; i++ {
		wg.Add(1)
		ps.jobs <- i
	}

	go func() {
		// ждём, пока все jobs будут сделаны
		wg.Wait()
		close(ps.jobs)
		close(ps.result)
	}()

	// 1. Закрыть result, только когда все worker точно отработали
	// 2. Когда worker точно отработали? Или отменён ctx, или канал jobs закрыт
	// 3. Когда можно канал jobs закрывать? Когда все jobs были обработаны

	var openPorts []int
	for p := range ps.result {
		openPorts = append(openPorts, p)
	}

	return openPorts
}

// worker starts worker
//
// about golang concurrency pipelines: https://go.dev/blog/pipelines
func (ps *PortScanner) worker(ctx context.Context, wg *sync.WaitGroup) {
	for {
		select {
		case port, ok := <-ps.jobs:
			fmt.Println("port", port, "ok", ok)

			if errors.Is(ctx.Err(), context.Canceled) {
				fmt.Println("stop doing job due to ctx canceled")
				return
			}

			if !ok {
				// worker отработал, когда канал jobs закрыт
				fmt.Println("jobs channel closed")
				return
			}

			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ps.host, port))
			if err != nil {
				// port closed
				wg.Done()
				continue
			}
			conn.Close()

			// port opened
			ps.result <- port
			fmt.Println("found open port", port)
			wg.Done()
		case <-ctx.Done():
			fmt.Println("context done")
			return
		}
	}
}
