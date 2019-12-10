package main

import (
	"github.com/gocolly/colly"
	"strings"
)

func after(value string, a string) string {
	// Get substring after a string.
	pos := strings.LastIndex(value, a)
	if pos == -1 {
		return ""
	}
	adjustedPos := pos + len(a)
	if adjustedPos >= len(value) {
		return ""
	}
	return value[adjustedPos:len(value)]
}
func unique(intSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range intSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
func extractSitemapRobots(website string) []string {
	e := colly.NewCollector(colly.AllowedDomains(website), colly.MaxDepth(2),
		colly.Async(true))

	sitemaps := []string{}

	e.OnResponse(func(r *colly.Response) {
		bodyString := string(r.Body)
		bodyStrings := strings.Split(bodyString, "\n")

		for _, st := range bodyStrings {
			if sitemap := after(st, "Sitemap:"); len(sitemap) > 0 {
				sitemaps = append(sitemaps, strings.TrimSpace(sitemap))

			}
		}

	})
	e.Visit("https://" + website + "/robots.txt")
	e.Wait()
	return sitemaps
}

func testSitemapURLS(website string) []string {
	site := []string{}
	e := colly.NewCollector(colly.Async(true))

	e.OnResponse(func(r *colly.Response) {

		if status := r.StatusCode; status == 200 {

			site = append(site, r.Request.URL.String())
		}
	})
	sitemapURLS := []string{"/sitemap.xml", "/sitemap_news.xml", "/sitemap_index.xml", "/sitemap-index.xml", "/sitemapindex.xml",
		"/sitemap-news.xml", "/post-sitemap.xml", "/page-sitemap.xml", "/portfolio-sitemap.xml", "home_slider-sitemap.xml", "category-sitemap.xml",
		"/author-sitemap.xml"}
	for i := range sitemapURLS {
		e.Visit("https://" + website + sitemapURLS[i])

	}

	e.Wait()
	return site
}

func getAllSitemap(website string) []string {
	robots := []string{}
	robots = extractSitemapRobots(website)
	sitemap := []string{}
	sitemap = testSitemapURLS(website)
	together := make([]string, len(robots)+len(sitemap))
	together = unique(append(together, append(sitemap, robots...)...))
	//fmt.Println("Found ", len(together), " sitemaps")
	//fmt.Println(together)
	return together
}
