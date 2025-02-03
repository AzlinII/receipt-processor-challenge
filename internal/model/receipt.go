package model

type Receipt struct {
	Retailer     string
	PurchaseDate string
	PurchaseTime string
	Items        []Item
	Total        string
}

type Item struct {
	ShortDescription string
	Price            string
}
