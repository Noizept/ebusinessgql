package scrapper

import (
	"sort"
	"strconv"
	"strings"

	"github.com/gocolly/colly"
)

// Citytel Signature struct for Citytel Website
type Citytel struct{}

// SearchItem Get Item results from Citytel search criteria
func (r *Citytel) SearchItem(itemName string) ItemSlice {
	var tmp ItemSlice

	c := colly.NewCollector(
		// MaxDepth is 2, so only the links on the scraped page
		// and links on those pages are visited
		colly.MaxDepth(2),
		colly.Async(true),
	)

	// Find and visit all links
	c.OnHTML(".article-box", func(e *colly.HTMLElement) {

		if e.ChildText(".title") != "" {
			hrefElem, _ := e.DOM.Find(".title").Attr("href")

			priceString := strings.ReplaceAll(e.ChildText(".price"), "лв.", "")
			priceString = strings.ReplaceAll(strings.ReplaceAll(priceString, ",", "."), " ", "")

			price, _ := strconv.ParseFloat(priceString, 32)

			imgURL, _ := e.DOM.Find(".first-image").Children().Html()
			imgURL = strings.ReplaceAll(imgURL, "/uploads", "https://www.citytel.bg/uploads")
			if price > 200 {
				tmp = append(tmp, Item{
					Name:     e.ChildText(".title"),
					Price:    price,
					URL:      e.Request.AbsoluteURL(hrefElem),
					ImageURL: imgURL,
					Website:  "citytel"})
			}
		}

	})

	c.Visit("https://www.citytel.bg/tarsi?search=" + strings.ReplaceAll(itemName, " ", "+"))
	c.Wait()
	sort.SliceStable(tmp, func(i, j int) bool {
		return tmp[i].Price > tmp[j].Price
	})
	return tmp
}
