package scrapper

// Item struct to hold extract data
type Item struct {
	Name     string  `json:"Name"`     // Name of the item
	Price    float64 `json:"Price"`    // Price of the item
	URL      string  `json:"URL"`      //Url of the item
	ImageURL string  `json:"ImageURL"` //Image Url of the item
	Website  string  `json:"Website"`
}

// ItemSlice  Dictionary with the list of Items
type ItemSlice []Item

// ItemDictionary List of items aggregated
type ItemDictionary map[string]ItemSlice
