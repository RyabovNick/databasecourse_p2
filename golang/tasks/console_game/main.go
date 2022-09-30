package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Enter command:")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	input := scanner.Text()
	fmt.Println("You entered:", input)
}
