package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"math"
)

type ProvideDiscount struct{}

func (ProvideDiscount) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	priceAfterDiscount := cart.TotalPrice - (cart.TotalPrice/100*promotion.DiscountPercentage)
	return dto.ShoppingCart{
		Products: cart.Products,
		TotalPrice: math.Round(priceAfterDiscount * 100)/100,
	}
}