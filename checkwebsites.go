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
			// send statement. send result struct for each call
			// to wc to the resultChannel
			// <- send operator
			// this is to prevent a race condition from goroutine writing to map at the same
			resultChannel <- result{u, wc(u)}
		}(url)
	}

	for i := 0; i < len(urls); i++ {
		// using a receive expression
		// which assigns a value received from a channel to a variable
		valueRecieved := <-resultChannel
		results[valueRecieved.string] = valueRecieved.bool
	}

	return results
}
