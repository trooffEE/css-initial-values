package main

import (
	"fmt"
	"github.com/gocolly/colly"
	"toffee.server.cssinitial/cmd/web/helpers"
)

type ScrappedCSSFormalDefinition struct {
	Initial map[string]string
}

func BuildHTMLQuery(path []string) string {
	return helpers.Join(" ", path...)
}

func getCssInitialDataByQuery(query string) (ScrappedCSSFormalDefinition, error) {
	c := colly.NewCollector(colly.Async(true))
	cssData := ScrappedCSSFormalDefinition{
		Initial: make(map[string]string),
	}

	htmlSelector := BuildHTMLQuery([]string{
		"[aria-labelledby='formal_definition']",
		".properties",
		"tr:contains('Initial value')",
		"td",
	})

	// Note: Данный скрапер нужен просто на случай, если у нас несколько значений initial
	// Это не кажется оптимальным, но пока что так 🤡
	c.OnHTML(fmt.Sprintf("%s %s", htmlSelector, "ul"), func(e *colly.HTMLElement) {
		if e.Text == "" {
			return
		}

		tempKey := ""
		e.ForEach("code", func(index int, child *colly.HTMLElement) {
			if (index+1)%2 != 0 {
				tempKey = child.Text
			} else {
				cssData.Initial[tempKey] = child.Text
			}
		})
	})

	c.OnHTML(htmlSelector, func(e *colly.HTMLElement) {
		if e.ChildText("ul") != "" {
			return
		}

		cssData.Initial[query] = e.Text
	})
	c.Wait()
	urlToScrap := fmt.Sprintf("%s%s", "https://developer.mozilla.org/en-US/docs/Web/CSS/", query)
	if err := c.Visit(urlToScrap); err != nil {
		return cssData, err
	}

	return cssData, nil
}
