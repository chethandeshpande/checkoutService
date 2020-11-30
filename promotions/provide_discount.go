package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideDiscount struct{}

func (ProvideDiscount) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	return dto.ShoppingCart{
		Products: cart.Products,
		TotalPrice: cart.TotalPrice - (cart.TotalPrice / 100 * promotion.DiscountPercentage),
	}
}