package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func itemCommand(itemName string) string {
	var result string
	itemName = strings.Replace(itemName, " ", "+", -1)

	crawler := colly.NewCollector()

	crawler.OnHTML("a[href]", func (e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 40 {
			if strings.Compare(link[0:39], "/url?q=https://classic.wowhead.com/item") == 0{
				if result == "" {
					result = itemDetails(link[7:])
				}
			}
		}
	})

	err := crawler.Visit("https://www.google.com/search?q=classic+wow+"+ itemName + "+wowhead")
	if err != nil {
		fmt.Printf("%v", err)
	}

	return result
}

func itemDetails(url string) string {
	var details string

	crawler := colly.NewCollector()

	crawler.OnHTML("meta[property='twitter:description']", func (e *colly.HTMLElement) {
		desc := e.Attr("content")

		details += desc
	})

	crawler.OnHTML("meta[property='twitter:image']", func (e *colly.HTMLElement) {
		image := e.Attr("content")

		details += "\n" + image
	})

	err := crawler.Visit(url)
	if err != nil {
		fmt.Printf("%v", err)
	}

	//fmt.Printf(details + "\n\n")
	return details
}
