package concurrency

type Checker func(string) bool

// TODO: Исправить Data race в коде
// 1. Mutex
// 2. Channel

func CheckWebsites(wc Checker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		go func(u string) {
			results[u] = wc(u)
		}(url)
	}

	return results
}

// url = a
// url = b
// url = c
// goroutine 1 work - c
// goroutine 2 work - c
// goroutine 3 work - c
