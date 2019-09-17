package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func dungeonCommand(dungeonName string) string {
	var result string
	dungeonName = strings.Replace(dungeonName, " ", "+", -1)

	crawler := colly.NewCollector()

	crawler.OnHTML("a[href]", func (e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 40 {
			if strings.Compare(link[0:35], "/url?q=https://classic.wowhead.com/") == 0{
				if result == "" {
					result = link[7:]
				}
			}
		}
	})

	err := crawler.Visit("https://www.google.com/search?q=classic+wow+"+ dungeonName + "+wowhead+strategy+guide")
	if err != nil {
		fmt.Printf("%v", err)
	}

	return result
}
