package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

func GetApplicablePromos(cart dto.ShoppingCart, promotions data.ProductPromotionMap) []data.PromotionDetails {
	alreadyAppliedPromos := make(map[string]bool)
	var applicablePromos []data.PromotionDetails
	for _, product := range cart.Products {
		if _, isPromoAvailableForProduct := promotions.Promotions[product.Name]; isPromoAvailableForProduct {
			if _, isPromoAlreadyApplied := alreadyAppliedPromos[product.Name]; !isPromoAlreadyApplied {
				promo := promotions.Promotions[product.Name]
				if isPromoEligible(promo, cart) {
					alreadyAppliedPromos[product.Name] = true
					applicablePromos = append(applicablePromos, promo)
				}
			}
		}
	}
	return applicablePromos
}

func isPromoEligible(promo data.PromotionDetails, cart dto.ShoppingCart) bool {
	var count int
	for _, product := range cart.Products {
		if product.Name == promo.ProductName {
			count++
		}
	}

	return count >= promo.MinimumQuantity
}
