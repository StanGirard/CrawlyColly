package main

import (
	"bufio"
	"fmt"
	"os"
	"sync"
)

func getSitemap(queue chan string, waitgroup *sync.WaitGroup) {
	
	website := <-queue
	waitgroup.Add(1)
	defer waitgroup.Done()
	crawled := crawlAll(website, getAllSitemap(website))
	uniqueCrawled := unique(crawled)
	message := write_to_file(uniqueCrawled, website, len(uniqueCrawled))
	fmt.Println(message)
	
	return 

}

func main() {

	website := make(chan string)
	var waitgroup sync.WaitGroup

	for i := 0; i < 10; i++ {
		go getSitemap(website, &waitgroup)
	}

	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		website <- line
	}
	
	waitgroup.Wait()
	fmt.Println("Done")

}
