package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

func ApplyPromos(cart dto.ShoppingCart, promos []data.PromotionDetails) dto.ShoppingCart {
	cartAfterAppliedPromos := cart
	for _, promo := range promos {
		promoStrategy := getPromoStrategy(promo.PromotionName)
		cartAfterAppliedPromos = promoStrategy.Apply(cartAfterAppliedPromos, promo)
	}
	return cartAfterAppliedPromos
}
