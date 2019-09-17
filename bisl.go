package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly"
	"strings"
)

func bislCommand(className string) string {
	className = strings.Replace(className, " ", "-", -1)

	var url string

	if len(className) < 8 {
		url = "https://www.icy-veins.com/wow-classic/" + className + "-dps-pve-gear-best-in-slot"
	} else {
		url = "https://www.icy-veins.com/wow-classic/" + className + "-pve-gear-best-in-slot"
	}

	var details string
	var bisListNames [20]string
	var bisListItems [20]string
	var bisListNamesCounter = 0
	var bisListItemsCounter = 0

	crawler := colly.NewCollector()

	crawler.OnHTML("body", func(body *colly.HTMLElement) {
		body.DOM.Find("table").First().Find("tr>td").Parent().Each(func(_ int, s *goquery.Selection) {
			symbol := s.Find("td").First().Text()
			bisListNames[bisListNamesCounter] = symbol
			bisListNamesCounter++
		})
	})

	crawler.OnHTML("body", func(body *colly.HTMLElement) {
		body.DOM.Find("table").First().Find("tr>td").Parent().Each(func(_ int, s *goquery.Selection) {
			symbol := s.Find("span").First().Text()
			bisListItems[bisListItemsCounter] = symbol
			bisListItemsCounter++
		})
	})

	err := crawler.Visit(url)
	if err != nil {
		fmt.Printf("Visiting icy veins error: %v", err)
	}

	for i := 0; i < 20; i++ {
		if bisListNames[i] != "" {
			details += bisListNames[i] + " ---> " + bisListItems[i] + "\n"
		}
	}

	return details
}
