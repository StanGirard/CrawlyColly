package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

func crawl(website string, urls string) []string {
	// Array containing all the known URLs in a sitemap
	knownUrls := []string{}

	c := colly.NewCollector(colly.Async(true))
	c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 20})
	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//url/loc", func(e *colly.XMLElement) {
		if len(e.Text) > 0 {
			knownUrls = append(knownUrls, strings.TrimSpace(e.Text))
		}

	})

	c.OnResponse(func(e *colly.Response) {
		fmt.Println(e.Request.URL)
	})
	c.OnXML("//sitemap/loc", func(e *colly.XMLElement) {
		c.Visit(e.Text)
	})
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 3 {
			if xxx := link[len(link)-3:]; xxx == "xml" {
				c.Visit(e.Request.AbsoluteURL(link))
			} else {
				knownUrls = append(knownUrls, e.Request.AbsoluteURL(link))
			}

		}

	})
	c.Visit(urls)
	c.Wait()

	return knownUrls

}

func crawlAll(domain string, sitemap []string) []string {
	known := []string{}
	for _, st := range sitemap {
		result := crawl(domain, st)
		println(st)
		println(len(result), ": Length")
		known = append(known, result...)
	}
	return known
}
