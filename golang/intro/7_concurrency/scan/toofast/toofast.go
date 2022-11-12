package toofast

import (
	"fmt"
	"net"
)

func Scan(address string) []int {
	var openPorts []int
	for i := 75; i < 100; i++ {
		go func(j int) {
			conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", j))
			if err != nil {
				fmt.Println(err)
				return
			}
			conn.Close()

			openPorts = append(openPorts, j)
			fmt.Println("Succesfull connection on port:", j)
		}(i)

		fmt.Println(i)
	}

	return openPorts
}
