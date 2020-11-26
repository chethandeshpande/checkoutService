package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

func getPotentialPromos(cart dto.ShoppingCart, promotions data.ProductPromotionMap) []data.PromotionDetails {
	keys := make(map[string]bool)
	var potentialPromos []data.PromotionDetails
	for _, product := range cart.Products {
		if _, value := keys[product.Name]; !value {
			keys[product.Name] = true
			potentialPromos = append(potentialPromos, promotions.Promotions[product.Name])
		}
	}
	return potentialPromos
}