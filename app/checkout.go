package app

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
)


func Checkout(cart dto.ShoppingCart, promotions data.ProductPromotionMap) dto.ShoppingCart {
	applicablePromos := services.GetPotentialPromos(cart, promotions)
	cartWithAppliedPromos := services.ApplyPromos(cart, applicablePromos)
	return cartWithAppliedPromos
}
