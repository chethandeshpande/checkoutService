package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideAFreeProduct struct{}

func (ProvideAFreeProduct) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	return dto.ShoppingCart{
		Products: append(cart.Products, dto.Product{
			Name:  promotion.AdditionalDetails.FreeProduct.Name,
			Price: promotion.AdditionalDetails.FreeProduct.Price,
		}),
		TotalPrice: cart.TotalPrice,
	}
}
