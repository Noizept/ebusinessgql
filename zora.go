package scrapper

import (
	"strings"

	"github.com/gocolly/colly"
)

// Zora Signature struct for Zora Website
type Zora struct{}

// SearchItem Get Item results from Technopolis search criteria
func (r *Zora) SearchItem(itemName string) ItemSlice {
	var tmp ItemSlice

	c := colly.NewCollector(
		// MaxDepth is 2, so only the links on the scraped page
		// and links on those pages are visited
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Find and visit all links
	c.OnHTML("._product-inner", func(e *colly.HTMLElement) {

		if e.ChildText("._product-name") != "" {
			hrefElem, _ := e.DOM.Find("._product-name").Children().Attr("href")
			//price, _ := strconv.ParseFloat(e.ChildText("_product-price"), 32)

			imgURL, _ := e.DOM.Find(".preview").Children().Html()

			tmp = append(tmp, Item{
				Name:     e.ChildText("._product-name"),
				Price:    000,
				URL:      e.Request.AbsoluteURL(hrefElem),
				ImageURL: imgURL,
				Website:  "techno"})

		}

	})

	c.Visit("https://zora.bg/products?search=" + strings.ReplaceAll(itemName, " ", "+") + "&order_by=price_from&order_direction=asc")
	c.Wait()
	return tmp
}
