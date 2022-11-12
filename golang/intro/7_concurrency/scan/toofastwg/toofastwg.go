package toofastwg

import (
	"fmt"
	"net"
	"sync"
)

func Scan(address string) []int {
	wg := sync.WaitGroup{}

	var openPorts []int
	for i := 1; i < 65535; i++ {
		wg.Add(1)
		go func(j int) {
			defer wg.Done()
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", address, j))
			if err != nil {
				// fmt.Println(err)
				return
			}
			conn.Close()

			openPorts = append(openPorts, j)
			fmt.Println("Successful connection on port:", j)
		}(i)
	}

	wg.Wait()

	return openPorts
}
