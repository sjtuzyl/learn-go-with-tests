package concurrency

type WebsiteChecker func(string) bool

type result struct {
	string
	bool
}

func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)
	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}
			// results[url] = wc(url)
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		ret := <-resultChannel
		results[ret.string] = ret.bool
	}
	return results
}
