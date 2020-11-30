package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideAFreeProduct struct{}

func (ProvideAFreeProduct) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	return dto.ShoppingCart{
		Products: append(cart.Products, dto.Product{
			Name: promotion.AdditionalDetails["FreeProduct"],
			Price: 30, //TODO: get Price of the product in additional details
		}),
		TotalPrice: cart.TotalPrice,
	}
}
