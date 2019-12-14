package scrapper

// Item struct to hold extract data
type Item struct {
	Name     string  // Name of the item
	Price    float64 // Price of the item
	URL      string  //Url of the item
	ImageURL string  //Image Url of the item
	Website  string
}

// ItemSlice  Dictionary with the list of Items
type ItemSlice []Item
