package dto

type ShoppingCart struct {
	Products   []Product
	TotalPrice float64
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}
