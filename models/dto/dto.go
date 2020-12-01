package dto

import "checkoutService/models/data"

type AdditionalDetails struct {
	AppliedPromos []data.PromotionDetails
}
type ShoppingCart struct {
	Products          []Product
	TotalPrice        float64
	AdditionalDetails AdditionalDetails
}

type Product struct {
	Name     string
	Price    float64
	Quantity int
}
