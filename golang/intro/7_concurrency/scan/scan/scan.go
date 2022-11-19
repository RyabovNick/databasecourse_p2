package scan

import (
	"fmt"
	"net"
	"sync"
)

// TODO: реализация в main не возвращает слайс открытых портов
// необходимо реализовать функцию Scan по по аналогии с кодом в main.go
// только её нужно дополнить, чтобы вернуть слайс открытых портов
// отсортированных по возрастанию

// 1. Если делаем slice с указанным cap. Передаём в worker index, таким образом каждый worker будет знать в какой index слайса писать, тем самым гонки не будет
// 		+ самый быстрый вариант реализации
// 		- не часто данный вариант подходит под задачу
// 2. Mutex, в worker передаётся указатель на Mutex, внутри worker перед append делается mu.Lock(), а после mu.Unlock.
// 		+ самая простая реализация
//		+ в некоторых задачах, где Lock потенциально очень малое количество, может быть быстрее коммуникации через channel
// 		- в большинстве случаев медленнее коммуникации через channel
// 3. Channel, все коммуникации между рутинами происходят через каналы
//		+ в большинстве случаев самый быстрый подход
// 		- реализация может быть сильно усложнена (для сравнения с реализацией через Mutex обязательно делать benchmark тесты)
// Также подход с worker может быть не самым оптимальным, каждую job можно запускать в отдельной goroutine, количество goroutine можно ограничить через channel или использовать https://pkg.go.dev/golang.org/x/sync/semaphore

func worker(ports <-chan int, wg *sync.WaitGroup) {
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

func Scan(address string) []int {
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

	// TODO: return result
	return nil
}
