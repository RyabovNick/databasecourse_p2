package slow

import (
	"fmt"
	"net"
)

func Scan(address string) []int {
	var openPorts []int
	for i := 75; i < 100; i++ {
		conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", "scanme.nmap.org", i))
		if err != nil {
			fmt.Println(err)
			continue
		}
		conn.Close()

		openPorts = append(openPorts, i)
		fmt.Println("Succesfull connection on port:", i)
	}

	return openPorts
}
