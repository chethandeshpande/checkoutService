package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"math"
)

type DiscountAmountOnTotalPrice struct{}

func (DiscountAmountOnTotalPrice) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	priceAfterDiscount := cart.TotalPrice - promotion.DiscountAmount
	return dto.ShoppingCart{
		Products:   cart.Products,
		TotalPrice: math.Round(priceAfterDiscount*100) / 100,
	}
}
