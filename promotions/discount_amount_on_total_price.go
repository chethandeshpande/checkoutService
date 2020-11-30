package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type DiscountAmountOnTotalPrice struct{}

func (DiscountAmountOnTotalPrice) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	return dto.ShoppingCart {
		Products:   cart.Products,
		TotalPrice: cart.TotalPrice - promotion.DiscountAmount,
	}
}
