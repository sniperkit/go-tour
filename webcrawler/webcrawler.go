package webcrawler

import (
	"fmt"
	"log"

	"github.com/sahilm/go-tour/concurrent"
)

// A web crawler.
type Crawler struct {
	url     string          // the url to crawl
	depth   int             // max-depth of links to follow
	fetcher Fetcher         // Fetcher to use
	visited *concurrent.Map // Map of visited urls
}

// Returns a new crawler.
func NewCrawler(url string, depth int, fetcher Fetcher) *Crawler {
	return &Crawler{
		url:     url,
		depth:   depth,
		fetcher: fetcher,
		visited: concurrent.NewMap(),
	}
}

// The Fetcher interface
type Fetcher interface {
	// Fetch returns a Page found at url
	Fetch(url string) (p *Page, err error)
}

// Represents a web page
type Page struct {
	Url   string   // The page's URL
	Body  string   // The body of the page
	Links []string // The links to other pages on the page
}

// The error returned when the URL cannot be found.
type ErrUrlNotFound string

// Error message of ErrUrlNotFound
func (e ErrUrlNotFound) Error() string {
	return fmt.Sprintf("%v not found", string(e))
}

// Crawl uses fetcher to recursively crawl
// pages starting with url, to a maximum of depth.
// Returns all pages found and errors collected.
func (c *Crawler) Crawl() ([]*Page, []error) {
	var pages []*Page
	var errors []error
	crawl(c.url, c.depth, c.fetcher, c.visited, &pages, &errors)
	return pages, errors
}

func crawl(url string, depth int, fetcher Fetcher, visited *concurrent.Map, pages *[]*Page, errors *[]error) {
	if depth <= 0 {
		log.Printf("<- Done with %v, depth 0.\n", url)
		return
	}

	_, ok := visited.PutIfAbsent(url, struct{}{})
	if !ok {
		log.Printf("<- Done with %v, already visited.\n", url)
		return
	}

	page, err := fetcher.Fetch(url)
	visited.Put(url, struct{}{})

	if err != nil {
		log.Printf("<- Error on %v: %v\n", url, err)
		*errors = append(*errors, err)
		return
	}

	log.Printf("Found: %s %q\n", page.Url, page.Body)
	*pages = append(*pages, page)
	done := make(chan bool)

	for i, l := range page.Links {
		log.Printf("-> Crawling child %v/%v of %v : %v.\n", i, len(page.Links), url, l)
		go func(url string) {
			crawl(url, depth-1, fetcher, visited, pages, errors)
			done <- true
		}(l)
	}

	for i := range page.Links {
		log.Printf("<- [%v] %v/%v Waiting for child %v.\n", url, i, len(page.Links), page.Links[i])
		<-done
	}

	log.Printf("<- Done with %v\n", url)
}
