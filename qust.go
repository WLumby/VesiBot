package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"strings"
)

func questCommand(questName string) string {
	var result string
	questName = strings.Replace(questName, " ", "+", -1)

	crawler := colly.NewCollector()

	crawler.OnHTML("a[href]", func (e *colly.HTMLElement) {
		link := e.Attr("href")
		if len(link) > 40 {
			if strings.Compare(link[0:34], "/url?q=https://classic.wowhead.com") == 0{
				if result == "" {
					result = link[7:]
				}
			}
		}
	})

	err := crawler.Visit("https://www.google.com/search?q=classic+wow+"+ questName + "+quest+wowhead")
	if err != nil {
		fmt.Printf("%v", err)
	}

	return result
}