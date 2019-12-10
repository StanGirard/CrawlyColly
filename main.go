package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sync"
	"time"
)

func timeTrack(start time.Time, name string) {
	elapsed := time.Since(start)
	log.Printf("%s took %s", name, elapsed)
}
func getSitemap(website string, waitgroup *sync.WaitGroup) int {

	defer waitgroup.Done()
	crawled := crawlAll(website, getAllSitemap(website))
	uniqueCrawled := unique(crawled)
	message, number := write_to_file(uniqueCrawled, website, len(uniqueCrawled))
	fmt.Println(message)

	return number

}

func main() {
	defer timeTrack(time.Now(), "Crawling")
	website := make(chan string)
	number := 0
	var waitgroup sync.WaitGroup

	for i := 0; i < 2; i++ {
		go func() {
			for {
				domain := <-website
				waitgroup.Add(1)
				number += getSitemap(domain, &waitgroup)
			}
		}()

	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		website <- line
	}

	waitgroup.Wait()

	fmt.Println("Found ", number, " URLS in the crawled websites")

	time.Sleep(2 * time.Second)
	return
}
