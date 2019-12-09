# CrawlyColly Fast Crawler

The purpose of this crawler is to get all the pages of a website very quickly

It uses the sitemaps of a website to discover the pages
The drawback is that if the pages aren't in the sitemap they won't be discovered.
However, it is a very fast & efficient way to get thousands of pages in seconds.

## Installation

- Install Golang

## Use the crawler

- Set the website your want to crawl in main.go  `website := "www.yourwebsite.com"`
- Run with `go run *.go`
- The pages will be writen in urls_<www.yourwebsite.com>.csv

## Disclaimer

Please be advised that even if this crawler doesn't visit every pages of a website, it can be very intensive for large websites.
Feel free to make pull request to improve the crawler



