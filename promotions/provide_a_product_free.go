package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type ProvideAFreeProduct struct{}

func (ProvideAFreeProduct) Apply(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	if !isFreeProductAlreadyPresentInCart(cart, promotion) {
		return cartWithAddedFreeProduct(cart, promotion)
	}
	return cartWithUpdatedQuantityOfFreeProduct(cart, promotion)
}

func cartWithAddedFreeProduct(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	return dto.ShoppingCart{
		Products: append(cart.Products, dto.Product{
			Name:     promotion.AdditionalDetails.FreeProduct.Name,
			Price:    promotion.AdditionalDetails.FreeProduct.Price,
			Quantity: promotion.AdditionalDetails.FreeProduct.Quantity,
		}),
		TotalPrice: cart.TotalPrice,
	}
}

func cartWithUpdatedQuantityOfFreeProduct(cart dto.ShoppingCart, promotion data.PromotionDetails) dto.ShoppingCart {
	var products []dto.Product
	for index := range cart.Products {
		product := cart.Products[index]
		if product.Name == promotion.AdditionalDetails.FreeProduct.Name {
			product = dto.Product{
				Name:     product.Name,
				Price:    product.Price,
				Quantity: product.Quantity + 1,
			}
		}
		products = append(products, product)
	}
	return dto.ShoppingCart{
		Products:          products,
		TotalPrice:        cart.TotalPrice,
		AdditionalDetails: cart.AdditionalDetails,
	}
}

func isFreeProductAlreadyPresentInCart(cart dto.ShoppingCart, promotion data.PromotionDetails) bool {
	for i := 0; i < len(cart.Products); i++ {
		product := cart.Products[i]
		if product.Name == promotion.AdditionalDetails.FreeProduct.Name {
			return true
		}
	}
	return false
}
