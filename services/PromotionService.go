package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
)

func applyPromos(cart dto.ShoppingCart, promos []data.PromotionDetails) dto.ShoppingCart {
	for _, promo := range promos {
		switch promo.PromotionName {
		case "BuyTwoGetOne":
			waiveOneProductPrice := promotions.WaiveOneProductPrice{}
			return waiveOneProductPrice.Apply(cart, promo)
		case "RaspberryPi":
			provideRaspberryPiFree := promotions.ProvideAFreeProduct{}
			return provideRaspberryPiFree.Apply(cart, promo)
		case "TenPercentOff":
			provideTenPercentOff := promotions.ProvideDiscount{}
			return provideTenPercentOff.Apply(cart, promo)
		default:
			return cart
		}
	}
	return cart
}