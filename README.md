# CrawlyColly Fast Crawler

The purpose of this crawler is to get all the pages of a website very quickly

It uses the sitemaps of a website to discover the pages
The drawback is that if the pages aren't in the sitemap they won't be discovered.
However, it is a very fast & efficient way to get thousands of pages in seconds.

## Installation

- Install Golang

## Use the crawler

- Create the folder `urls` with the command `mkdir urls`
- Set the websites your want to crawl in a file like `urls_test`, one url per line
- Compile with `go build *.go`
- Run with `cat urls_test | ./crawl` or if not compiled `cat urls_test | go run *.go`
- The websites' urls will be writen in urls_<www.yourwebsite.com>.csv

At the end of your crawling if you want to merge all the files just run in the folder urls `for filename in $(ls *.csv); do sed 1d $filename >> ../final.csv; done`

## Disclaimer

Please be advised that even if this crawler doesn't visit every pages of a website, it can be very intensive for large websites.
Feel free to make pull request to improve the crawler



