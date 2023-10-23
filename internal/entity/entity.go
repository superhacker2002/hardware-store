package entity

type ShelfProduct struct {
	ShelfName         string
	ProductName       string
	ProductID         int
	OrderID           int
	Quantity          int
	AdditionalShelves []string
}
