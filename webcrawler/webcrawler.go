package webcrawler

import (
	"fmt"

	"log"

	"github.com/sahilm/go-tour/concurrent"
)

type Crawler struct {
	Url     string
	depth   int
	fetcher Fetcher
	visited *concurrent.Map
}

func NewCrawler(url string, depth int, fetcher Fetcher) *Crawler {
	return &Crawler{
		Url:     url,
		depth:   depth,
		fetcher: fetcher,
		visited: concurrent.NewMap(),
	}
}

type Fetcher interface {
	Fetch(url string) (p *Page, err error)
}

type Page struct {
	Url   string
	Body  string
	Links []string
}

type ErrUrlNotFound string

func (e ErrUrlNotFound) Error() string {
	return fmt.Sprintf("%v not found", string(e))
}

func (c *Crawler) Crawl() ([]*Page, []error) {
	var pages []*Page
	var errors []error
	crawl(c.Url, c.depth, c.fetcher, c.visited, &pages, &errors)
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
