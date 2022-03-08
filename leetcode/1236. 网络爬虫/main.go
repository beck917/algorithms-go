package main

import (
	"fmt"
	"sync"
)

func main() {
	sim := &SimParser{}

	results := crawl("http://news.yahoo.com", sim)

	fmt.Println(results)
}

type HtmlParser interface {
	getUrls(url string) []string
}

type SimParser struct {
}

func (sim *SimParser) getUrls(url string) []string {
	if url == "http://news.yahoo.com" {
		return []string{
			"http://news.yahoo.com/news",
			"http://news.yahoo.com/news/topics/",
		}
	} else if url == "http://news.yahoo.com/news" {
		return []string{
			"http://news.yahoo.com/us",
			"http://news.yahoo.com",
			"http://news.yahoo.com/news/topics/",
		}
	}

	return nil
}

var visited map[string]bool
var results []string
var mu sync.RWMutex
var wg sync.WaitGroup

func crawl(startUrl string, htmlParser HtmlParser) []string {
	visited = make(map[string]bool)
	results = make([]string, 0)

	dfs(startUrl, htmlParser)

	return results
}

func dfs(startUrl string, htmlParser HtmlParser) {
	mu.Lock()
	if visited[startUrl] == true {
		mu.Unlock()
		return
	}
	visited[startUrl] = true
	results = append(results, startUrl)
	mu.Unlock()

	urls := htmlParser.getUrls(startUrl)

	var tmpwg sync.WaitGroup
	for _, url := range urls {
		tmpwg.Add(1)
		go func(url string) {
			dfs(url, htmlParser)
			tmpwg.Done()
		}(url)
	}
	tmpwg.Wait()
}
