package scrapper

// TODO: REMAKE ALL THE PACKAGE

import (
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Emag Signature struct for Emag Website
type Emag struct{}

// SearchItem Get Item results from Emag search criteria
func (r *Emag) SearchItem(itemName string) ItemSlice {
	var tmp ItemSlice

	c := colly.NewCollector(
		// MaxDepth is 2, so only the links on the scraped page
		// and links on those pages are visited
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Find and visit all links
	c.OnHTML(".js-product-data", func(e *colly.HTMLElement) {

		if e.ChildText("._product-name") != "" {
			hrefElem, _ := e.DOM.Find("._product-name").Children().Children().Attr("href")
			priceString := strings.ReplaceAll(e.ChildTexts("._product-price")[0], "лв.", "")
			priceString = strings.TrimSpace(strings.ReplaceAll(priceString, ",", "."))

			price, _ := strconv.ParseFloat(priceString, 32)

			imgURL, _ := e.DOM.Find("._product-image-thumb-holder").Html()
			tmp = append(tmp, Item{
				Name:     e.ChildText("._product-name"),
				Price:    price,
				URL:      e.Request.AbsoluteURL(hrefElem),
				ImageURL: imgURL,
				Website:  "emag"})

		}

	})

	c.Visit("https://www.emag.bg/search/" + strings.ReplaceAll(itemName, " ", "%20") + "?ref=effective_search")
	c.Wait()
	return tmp
}
