package main

import (
	"bufio"
	"fmt"

	"os"
	"time"
)

func getSitemap(queue chan string) {

	website := <-queue
	crawled := crawlAll(website, getAllSitemap(website))
	uniqueCrawled := unique(crawled)
	message := write_to_file(uniqueCrawled, website, len(uniqueCrawled))
	fmt.Println(message)

}

func main() {

	website := make(chan string)
	for i := 0; i < 4; i++ {
		go getSitemap(website)
	}
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		
		website <- line
	}
	time.Sleep(10 * time.Second)

}
