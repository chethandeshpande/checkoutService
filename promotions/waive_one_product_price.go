package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type WaiveOneProductPrice struct{}

func (WaiveOneProductPrice) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
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
		TotalPrice: cart.TotalPrice - (cart.Products[0].Price),
	}
}
