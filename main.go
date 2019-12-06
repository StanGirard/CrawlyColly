package main

import (
	"fmt"
	"strings"

	"github.com/gocolly/colly"
)

/*
func main() {
	// Array containing all the known URLs in a sitemap
	website := "www.leparisien.fr"
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains(website), colly.MaxDepth(2),
		colly.Async(true))
	//c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {

		if len(e.Text) > 0 {
			knownUrls = append(knownUrls, strings.TrimSpace(e.Text))
		}

	})

	c.OnXML("//sitemapindex/sitemap/loc", func(e *colly.XMLElement) {
		c.Visit(e.Text)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		if len(link) > 0 {
			knownUrls = append(knownUrls, e.Request.AbsoluteURL(link))
			c.Visit(e.Request.AbsoluteURL(link))
		}

	})

	// Start the collector
	c.Visit("https://" + website + "/sitemap.xml")
	c.Wait()
	if len(knownUrls) == 0 {
		fmt.Println("Not URLS found at : https://" + website + "/sitemap.xml")
		c.Visit("https://" + website + "/sitemap_index.xml")
	}
	c.Wait()
	if len(knownUrls) == 0 {
		fmt.Println("Not URLS found at : https://" + website + "/sitemap_index.xml")
		c.Visit("https://" + website + "/sitemap/")
	}
	c.Wait()
	if len(knownUrls) == 0 {
		fmt.Println("Not URLS found at : https://" + website + "/sitemap/")
		c.Visit("https://" + website)
	}
	c.Wait()
	fmt.Println("Collected", len(knownUrls), "URLs")
	write_to_file(knownUrls, website)
}*/

func main() {
	// Array containing all the known URLs in a sitemap
	website := "www.linternaute.com"
	knownUrls := []string{}

	// Create a Collector specifically for Shopify
	c := colly.NewCollector(colly.AllowedDomains(website), colly.MaxDepth(2),
		colly.Async(true))
	e := colly.NewCollector(colly.AllowedDomains(website), colly.MaxDepth(2), colly.Async(true))

	e.OnHTML("/html/body", func(r *colly.HTMLElement) {
		fmt.Println(r.Text)

	})
	e.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	//c.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 100})
	// Create a callback on the XPath query searching for the URLs
	c.OnXML("//urlset/url/loc", func(e *colly.XMLElement) {

		if len(e.Text) > 0 {
			knownUrls = append(knownUrls, strings.TrimSpace(e.Text))
		}

	})

	c.OnXML("//sitemapindex/sitemap/loc", func(e *colly.XMLElement) {
		c.Visit(e.Text)
	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {

		link := e.Attr("href")
		if len(link) > 0 {
			knownUrls = append(knownUrls, e.Request.AbsoluteURL(link))
			c.Visit(e.Request.AbsoluteURL(link))
		}

	})

	// Start the collector

	fmt.Println(getAllSitemap(website))
	//write_to_file(knownUrls, website)
}
