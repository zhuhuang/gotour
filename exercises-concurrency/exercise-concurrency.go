package main

import (
	"fmt"
	"sync"
)

type Fetcher interface {
	// Fetch returns the body of URL and
	// a slice of URLs found on that page.
	Fetch(url string) (body string, urls []string, err error)
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
func Crawl(url string, depth int, fetcher Fetcher, mutex *FetchMetux, ch chan CrawlerResponse) {
	// TODO: Fetch URLs in parallel.
	// TODO: Don't fetch the same URL twice.
	// This implementation doesn't do either:
	defer close(ch)

	if depth <= 0 {
		return
	}

	// duplicate job check
	mutex.mux.Lock()
	if mutex.flags[url] {
		mutex.mux.Unlock()
		return
	}
	mutex.flags[url] = true
	mutex.mux.Unlock()

	body, urls, err := fetcher.Fetch(url)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("found: %s %q\n", url, body)
	ch <- CrawlerResponse{url, body}

	resultChannels := make([]chan CrawlerResponse, len(urls))
	for i, u := range urls {
		resultChannels[i] = make(chan CrawlerResponse)
		go Crawl(u, depth-1, fetcher, mutex, resultChannels[i])
	}

	for i := range resultChannels {
		for response := range resultChannels[i] {
			ch <- response
		}
	}

	return
}

func main() {
	mutex := &FetchMetux{flags: make(map[string]bool)}
	resultChannel := make(chan CrawlerResponse)

	go Crawl("https://golang.org/", 4, fetcher, mutex, resultChannel)

	for response := range resultChannel {
		fmt.Printf("found: %s %q\n", response.url, response.body)
	}
}

// fakeFetcher is Fetcher that returns canned results.
type FakeFetcher map[string]*FakeResult

type FakeResult struct {
	body string
	urls []string
}

func (f FakeFetcher) Fetch(url string) (string, []string, error) {
	if res, ok := f[url]; ok {
		return res.body, res.urls, nil
	}
	return "", nil, fmt.Errorf("not found: %s", url)
}

// fetcher is a populated fakeFetcher.
var fetcher = FakeFetcher{
	"https://golang.org/": &FakeResult{
		"The Go Programming Language",
		[]string{
			"https://golang.org/pkg/",
			"https://golang.org/cmd/",
		},
	},
	"https://golang.org/pkg/": &FakeResult{
		"Packages",
		[]string{
			"https://golang.org/",
			"https://golang.org/cmd/",
			"https://golang.org/pkg/fmt/",
			"https://golang.org/pkg/os/",
		},
	},
	"https://golang.org/pkg/fmt/": &FakeResult{
		"Package fmt",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
	"https://golang.org/pkg/os/": &FakeResult{
		"Package os",
		[]string{
			"https://golang.org/",
			"https://golang.org/pkg/",
		},
	},
}

type FetchMetux struct {
	flags map[string]bool
	mux sync.Mutex
}

type CrawlerResponse struct {
	url string
	body string
}
