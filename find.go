package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func findCommand(npcName string) string {
	var result string
	npcName = strings.Replace(npcName, " ", "+", -1)

	crawler := colly.NewCollector()

	crawler.OnHTML("a[href]", func (e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 40 {
			if strings.Compare(link[0:38], "/url?q=https://classic.wowhead.com/npc") == 0{
				if result == "" {
					result = findDetails(link[7:])
				}
			}
		}
	})

	err := crawler.Visit("https://www.google.com/search?q=classic+wow+"+ npcName + "+wowhead")
	if err != nil {
		fmt.Printf("%v", err)
	}

	return result
}

func findDetails(url string) string {
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
