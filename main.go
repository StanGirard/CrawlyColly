package main

import "fmt"

func main() {
	website := "www.liberation.fr"
	crawled := crawlAll(website, getAllSitemap(website))
	uniqueCrawled := unique(crawled)
	fmt.Println("Visited ", len(uniqueCrawled), " pages on ", website)
	write_to_file(crawled, website)
}
