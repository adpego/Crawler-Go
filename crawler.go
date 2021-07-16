package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"sort"
	"strings"
)

func find(slice []string, elemFind string) bool {
	for _, elem := range slice {
		if elem == elemFind {
			return true
		}
	}
	return false
}

func main() {

	var urls []string
	var images []string
	var scripts []string
	var links []string

	c := colly.NewCollector(
		colly.MaxDepth(1),
		colly.AllowedDomains("adpego.com"),
	)

	// Find and visit all links
	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		if !strings.HasPrefix(e.Attr("href"), "#") {
			url := e.Request.AbsoluteURL(e.Attr("href"))
			found := find(urls, url)
			if !found {
				urls = append(urls, url)
			}
			fmt.Println("VISITING", e.Attr("href"))
			e.Request.Visit(e.Attr("href"))
		}

	})

	c.OnHTML("img[src]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))
		found := find(images, url)
		if !string.HasPrefix(e.Attr("src"), "data:image") {
			if !found {
				images = append(images, url)
			}
		]

	})

	c.OnHTML("link[href]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))
		found := find(links, url)
		if !found {
			links = append(links, url)
		}
	})

	c.OnHTML("script[src]", func(e *colly.HTMLElement) {
		url := e.Request.AbsoluteURL(e.Attr("src"))
		found := find(scripts, url)
		if !found {
			scripts = append(scripts, url)
		}
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
