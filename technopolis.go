package scrapper

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Technopolis Signature struct for Technopolis Website
type Technopolis struct{}

// SearchItem Get Item results from Technopolis search criteria
func (r *Technopolis) SearchItem(itemName string) ItemSlice {
	var tmp  ItemSlice

	c := colly.NewCollector(
		// MaxDepth is 2, so only the links on the scraped page
		// and links on those pages are visited
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Find and visit all links
	c.OnHTML(".list-item", func(e *colly.HTMLElement) {

		if e.ChildText(".item-name") != "" {
			hrefElem, _ := e.DOM.Find(".item-name").Children().Attr("href")
			price, _ := strconv.ParseFloat(e.ChildText(".price-value"), 32)

			imgURL, _ := e.DOM.Find(".preview").Children().Html()

			tmp = append(tmp, Item{
				Name:     e.ChildText(".item-name"),
				Price:    price,
				URL:      e.Request.AbsoluteURL(hrefElem),
				ImageURL: imgURL,
				Website:  "techno"})

		}

	})

	c.Visit("https://www.technopolis.bg/bg/search/?query=" + strings.ReplaceAll(itemName, " ", "%20"))
	c.Wait()
	return tmp
}
