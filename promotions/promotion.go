package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type Promotion interface {
	Apply(cart dto.ShoppingCart, details data.PromotionDetails) dto.ShoppingCart
}