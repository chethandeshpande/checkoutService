package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideAFreeProduct struct{}

func (ProvideAFreeProduct) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	var count int
	for _, product := range cart.Products {
		if product.Name == promotion.ProductName {
			count++
		}
	}
	if count < promotion.MinimumQuantity {
		return cart
	}

	return dto.ShoppingCart{
		Products: append(cart.Products, dto.Product{
			Name: promotion.AdditionalDetails["FreeProduct"],
			Price: 30, //TODO: get Price of the product in additional details
		}),
		TotalPrice: cart.TotalPrice,
	}
}
