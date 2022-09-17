package main

import "fmt"

func main() {
	// бесконечный цикл
	// for {
	// 	continue

	// 	break
	// }

	// for a < b {
	// 	...
	// }

	// for i := 0; i < 10; i++ {
	// 	...
	// }

	for _, v := range []string{"1", "2"} {
		fmt.Println(v)
	}

	// &v - всегда одинаковый адрес, т.к. v переиспользуется в range цикле
	in := []string{"1", "2", "3"}
	for _, v := range in {
		fmt.Printf("address of %v: %v\n", v, &v)
	}
}
