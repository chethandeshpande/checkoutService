package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type DiscountAmountOnTotalPrice struct{}

func (DiscountAmountOnTotalPrice) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	var count int
	for _, product := range cart.Products {
		if product.Name == promotion.ProductName {
			count++
		}
	}
	if count < promotion.MinimumQuantity {
		return cart
	}

	return dto.ShoppingCart {
		Products:   cart.Products,
		TotalPrice: cart.TotalPrice - promotion.DiscountAmount,
	}
}
