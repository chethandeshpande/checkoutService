package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideDiscount struct{}

func (ProvideDiscount) Apply(cart dto.ShoppingCart, details data.PromotionDetails) dto.ShoppingCart {
	var count int
	for _, product := range cart.Products {
		if product.Name == details.ProductName {
			count++
		}
	}
	if count < details.MinimumQuantity {
		return cart
	}
	return dto.ShoppingCart{
		Products: cart.Products,
		TotalPrice: cart.TotalPrice - (cart.TotalPrice / 100 * details.DiscountPercentage),
	}
}