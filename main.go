package main

import (
	"checkoutService/app"
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

func main() {
	macbookPro := dto.Product{
		Name:  "Macbook Pro",
		Price: 1,
	}
	//alexaSpeaker := dto.Product{
	//	Name:  "Alexa Speaker",
	//	Price: 1,
	//}
	//googleHome := dto.Product{
	//	Name:  "Google Home",
	//	Price: 1,
	//}
	//products := []dto.Product{macbookPro, alexaSpeaker, googleHome}

	macbookProPromo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaSpeakerPromo := data.PromotionDetails{
		PromotionName:      "10PercentOff",
		ProductName:        "Alexa Speaker",
		MinimumQuantity:    3,
		DiscountPercentage: 10,
	}
	googleHomePromo := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    2,
		DiscountPercentage: 0,
	}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = macbookProPromo
	promoToProductMap["Alexa Speaker"] = alexaSpeakerPromo
	promoToProductMap["Google Home"] = googleHomePromo

	shoppingCart := dto.ShoppingCart{
		Products:   []dto.Product{macbookPro},
		TotalPrice: 1,
	}

	cartAfterAppliedPromos := app.Checkout(shoppingCart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	for _, product := range cartAfterAppliedPromos.Products {
		println(product.Name)
		println(product.Price)
	}
	println(cartAfterAppliedPromos.TotalPrice)
}