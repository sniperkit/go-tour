// Adapted from the official solution with the following copyright: Copyright 2012 The Go Authors.  All rights reserved.
package webcrawler

import (
	"fmt"

	"github.com/sahilm/go-tour/concurrent"
)

// Crawler is a web crawler
type Crawler struct {
	url     string          // the url to crawl
	depth   int             // max-depth of links to follow
	fetcher Fetcher         // Fetcher to use
	visited *concurrent.Map // Map of visited urls
}

// NewCrawler returns a new crawler.
func NewCrawler(url string, depth int, fetcher Fetcher) *Crawler {
	return &Crawler{
		url:     url,
		depth:   depth,
		fetcher: fetcher,
		visited: concurrent.NewMap(),
	}
}

// Fetcher can fetch web pages
type Fetcher interface {
	// Fetch returns a Page found at url
	Fetch(url string) (p *Page, err error)
}

// Page represents a web page
type Page struct {
	URL   string   // The page's URL
	Body  string   // The body of the page
	Links []string // The links to other pages on the page
}

// ErrURLNotFound is the error representing a page that cannot be loaded.
type ErrURLNotFound string

// Error message of ErrURLNotFound
func (e ErrURLNotFound) Error() string {
	return fmt.Sprintf("%v not found", string(e))
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Returns all pages found and errors collected.
func (c *Crawler) Crawl() ([]*Page, []error) {
	p := concurrent.NewSlice()
	e := concurrent.NewSlice()

	crawl(c.url, c.depth, c.fetcher, c.visited, p, e)

	pView := p.View()
	var pages []*Page
	for i := 0; i < len(pView); i++ {
		pages = append(pages, pView[i].(*Page))
	}

	eView := e.View()
	var errors []error
	for i := 0; i < len(eView); i++ {
		errors = append(errors, eView[i].(error))
	}
	return pages, errors
}

var sentinel = struct{}{}

func crawl(url string, depth int, fetcher Fetcher, visited *concurrent.Map, pages *concurrent.Slice, errors *concurrent.Slice) {
	if depth <= 0 {
		return
	}

	_, ok := visited.PutIfAbsent(url, sentinel)
	if !ok {
		return
	}

	page, err := fetcher.Fetch(url)
	visited.Put(url, sentinel)

	if err != nil {
		errors.Append(err)
		return
	}

	pages.Append(page)
	done := make(chan bool)

	for _, url := range page.Links {
		go func(url string) {
			crawl(url, depth-1, fetcher, visited, pages, errors)
			done <- true
		}(url)
	}

	for range page.Links {
		<-done
	}
}
