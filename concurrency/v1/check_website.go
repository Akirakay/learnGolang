package concurrency

//WebsiteChecker is a type of func
type WebsiteChecker func(string) bool

//CheckWebsites return slice of checked urls
func CheckWebsites(wc WebsiteChecker, urls []string) map[string]bool {
	results := make(map[string]bool)

	for _, url := range urls {
		results[url] = wc(url)
	}

	return results
}
