package summator

import (
	"fmt"
	"sync"
)

// Задача:
// Даётся на вход N чисел
// Запустить в количесте рутин из расчёта N/M чисел
const M = 4

func Summator(input []int) int {
	sumch := make(chan int, len(input)/M)
	summ := 0
	wg := sync.WaitGroup{}

	for i := 0; i < len(input)/M; i++ {
		wg.Add(1)
		go func(in []int) {
			defer wg.Done()
			fmt.Println(in)
			sum := 0
			for _, v := range in {
				sum += v
			}
			sumch <- sum
		}(input[i*M : (i+1)*M])
	}

	wg.Wait()
	close(sumch)
	// go func() {
	// 	wg.Wait()
	// 	close(sumch)
	// }()

	for i := range sumch {
		summ += i
		fmt.Println("get", i)
	}

	return summ
}
