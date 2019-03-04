package engine

var visitedUrls = make(map[string]bool)

// Web Page to repeat
func isDuplicate(url string) bool {
	if visitedUrls[url] {
		// log.Error("repeat: %s", url)
		return true
	}
	visitedUrls[url] = true

	return false
}
