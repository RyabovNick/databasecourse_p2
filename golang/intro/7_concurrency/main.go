package main

import (
	"fmt"
	"net"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup) {
	for p := range ports {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "127.0.0.1", p))
		if err != nil {
			// fmt.Println("port closed:", p)
			wg.Done()
			continue
		}
		conn.Close()

		fmt.Println("port opened:", p)

		wg.Done()
	}
}

func main() {
	ports := make(chan int, 200)
	wg := sync.WaitGroup{}

	for i := 0; i < cap(ports); i++ {
		go worker(ports, &wg)
	}

	for i := 1; i < 10000; i++ {
		wg.Add(1)
		ports <- i
	}

	wg.Wait()
	close(ports) // канал закрывает обязательно тот, кто в него пишет
	// иначе возможна паника
}
