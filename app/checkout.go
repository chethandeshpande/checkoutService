package app

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
)

func Checkout(cart dto.ShoppingCart, promotions data.ProductPromotionMap) dto.ShoppingCart {
	applicablePromos := services.GetApplicablePromos(cart, promotions)
	cartWithAppliedPromos := services.ApplyPromos(cart, applicablePromos)
	cartWithAppliedPromos.AdditionalDetails = dto.AdditionalDetails{
		AppliedPromos: applicablePromos,
	}
	return cartWithAppliedPromos
}
