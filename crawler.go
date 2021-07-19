package main

import (
	"flag"
	"fmt"
	"sort"
	"strings"

	"github.com/gocolly/colly"
)

var (
	word              string
	maxDepth          int
	verbose           bool
	noOutput          bool
	allAllowedDomains bool
	allowedDomains    string
	filename          string
)

func find(slice []string, elemFind string) bool {
	for _, elem := range slice {
		if elem == elemFind {
			return true
		}
	}
	return false
}

func appendIfNotExist(slice []string, elem string) []string {
	if !find(slice, elem) {
		slice = append(slice, elem)
	}
	return slice
}

func init() {
	flag.StringVar(&word, "u", "", "Url to start the crawler")
	flag.IntVar(&maxDepth, "d", 3, "Max depth for crawling")
	flag.BoolVar(&verbose, "v", false, "Verbose option, show links visiting")
	flag.BoolVar(&noOutput, "no", false, "No output - option, for don't show the result to command line, requirement -f parameter")
	flag.StringVar(&filename, "f", "", "File to save the output")
	flag.StringVar(&allowedDomains, "ad", "", "Allowed domains to crawl. (Default value are the url)")
	flag.BoolVar(&allAllowedDomains, "all", false, "Allow all domains to crawl, not recomended for a uncontrolled environment")

}

func main() {

	flag.Parse()

	var urls []string
	var images []string
	var scripts []string
	var links []string

	/*c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.AllowedDomains("adpego.com"),
	)*/
	c := colly.NewCollector()
	c.MaxDepth = 1
	colly.AllowedDomains("adpego.com")

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if !strings.HasPrefix(e.Attr("href"), "#") {
			url := e.Request.AbsoluteURL(e.Attr("href"))

			urls = appendIfNotExist(urls, url)

			fmt.Println("VISITING", e.Attr("href"))
			e.Request.Visit(e.Attr("href"))
		}

	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))

		if !strings.HasPrefix(e.Attr("src"), "data:image") {
			images = appendIfNotExist(images, url)
		}

	})

	c.OnHTML("link[href]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))
		links = appendIfNotExist(links, url)

	})

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))
		scripts = appendIfNotExist(scripts, url)

	})

	c.Visit("https://adpego.com")

	sort.Strings(urls)
	sort.Strings(scripts)
	sort.Strings(images)
	sort.Strings(links)

	fmt.Println("====== URL ======")
	for _, url := range urls {
		fmt.Println(url)
	}

	fmt.Println("\n====== SCRIPTS ======")
	for _, script := range scripts {
		fmt.Println(script)
	}

	fmt.Println("\n====== IMG ======")
	for _, img := range images {
		fmt.Println(img)
	}

	fmt.Println("\n====== LINKS ======")
	for _, link := range links {
		fmt.Println(link)
	}

}
