package scrapper

// Item struct to hold extract data
type Item struct {
	Name     string  // Name of the item
	Price    float64 // Price of the item
	URL      string  //Url of the item
	ImageURL string  //Image Url of the item
	Website  string
}

// ItemDictionary  Dictionary with the list of Items
type ItemDictionary map[string][]Item
