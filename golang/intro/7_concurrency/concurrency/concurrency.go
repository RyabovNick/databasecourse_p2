package concurrency

import "sync"

type Checker func(string) bool

// TODO: Исправить Data race в коде
// 1. Mutex + return раньше, чем все routine запишут значение
// 2. Channel (Homework)

func CheckWebsites(wc Checker, urls []string) map[string]bool {
	results := make(map[string]bool)
	mu := sync.Mutex{}
	wg := sync.WaitGroup{}

	for _, url := range urls {
		wg.Add(1)
		go func(u string) {
			defer wg.Done()
			mu.Lock()
			results[u] = wc(u)
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	return results
}

// url = a
// url = b
// url = c
// goroutine 1 work - c
// goroutine 2 work - c
// goroutine 3 work - c
