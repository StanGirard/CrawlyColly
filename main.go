package main

import (
	"fmt"
)

func main() {
	website := "www.lemonde.fr"
	crawled := crawlAll(website, getAllSitemap(website))
	uniqueCrawled :=  unique(crawled)
	fmt.Println("Visited ", len(uniqueCrawled), " pages on ", website)
	//crawled := crawl(website, "https://www.lemonde.fr/sitemap_index.xml")
	
	write_to_file(crawled, website)
}
