package main

import (
	"github.com/gocolly/colly"
)

type ScrappedCSSFormalDefinition struct {
	Initial map[string]string
}

// TODO вынести в отдельный пакет
func Join(separator string, path ...string) string {
	result := ""
	for i := range path {
		result += path[i]
		if i != len(path)-1 {
			result += separator
		}
	}
	return result
}

func buildHtmlQuery(path []string) string {
	return Join(" ", path...)
}

func getCssInitialDataByQuery(query string) (ScrappedCSSFormalDefinition, error) {
	cssData := ScrappedCSSFormalDefinition{
		Initial: make(map[string]string),
	}

	c := colly.NewCollector(colly.Async(true))

	//c.OnRequest(func(r *colly.Request) {
	//	fmt.Println("Visiting", r.URL)
	//})

	queryToSearch := []string{
		"[aria-labelledby='formal_definition']",
		".properties",
		"tr:contains('Initial value')",
		"td",
	}

	// Note: Данный скрапер нужен просто на случай, если у нас несколько значений initial
	// Это не кажется оптимальным, но пока что так 🤡
	c.OnHTML(buildHtmlQuery(queryToSearch)+" ul", func(e *colly.HTMLElement) {
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

	c.OnHTML(buildHtmlQuery(queryToSearch), func(e *colly.HTMLElement) {
		if e.ChildText("ul") != "" {
			return
		}

		cssData.Initial[query] = e.Text
	})

	err := c.Visit("https://developer.mozilla.org/en-US/docs/Web/CSS/" + query)
	c.Wait()
	if err != nil {
		return cssData, err
	}

	return cssData, nil
}
