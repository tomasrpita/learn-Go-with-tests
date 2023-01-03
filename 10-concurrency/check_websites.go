package concurrency

// WebsiteChequer cheacks a url, returning a bool.
type Websitechequer func(string) bool
type result struct {
	string
	bool
}

// CheckWebsites takes a WebsiteChequer and a slice of url and returns a map
// of urls to the result of checking each url with the  WebsiteChecker funtion.
func CheckWebsites(wc Websitechequer, urls []string) map[string]bool {
	results := make(map[string]bool)
	resultChannel := make(chan result)

	for _, url := range urls {
		go func(u string) {
			resultChannel <- result{u, wc(u)}

		}(url)
	}

	for i := 0; i < len(urls); i++ {
		r := <-resultChannel
		results[r.string] = r.bool
	}
	return results
}
