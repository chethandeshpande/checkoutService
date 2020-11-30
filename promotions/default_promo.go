package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type DefaultPromotion struct{}

func (DefaultPromotion) Apply(cart dto.ShoppingCart, details data.PromotionDetails) dto.ShoppingCart {
	return cart
}
