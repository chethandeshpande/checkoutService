package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

func GetPotentialPromos(cart dto.ShoppingCart, promotions data.ProductPromotionMap) []data.PromotionDetails {
	keys := make(map[string]bool)
	var potentialPromos []data.PromotionDetails
	for _, product := range cart.Products {
		if _, value := keys[product.Name]; !value {
			if _, promoPresent := promotions.Promotions[product.Name] ; promoPresent {
				keys[product.Name] = true
				potentialPromos = append(potentialPromos, promotions.Promotions[product.Name])

			}
		}
	}
	return potentialPromos
}