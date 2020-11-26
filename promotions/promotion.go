package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
)

type Promotion interface {
	Apply(cart dto.ShoppingCart, details data.PromotionDetails) dto.ShoppingCart
}


// Product - Name,
// Promotion - map of promotion to combination of products
// Cart - Products
// Inventory - map of product  quantity, Price

/*
Promotions: {
	"googleHome": {
			"quantity": 3
			"promo": "BuyTwoGetOne"
	},
	"macbookpro": {
		"quantity": 1,
		"promo": "AddGiftProduct"
		"giftProduct": "RaspberryPi"
	},
	"alexa": {
		"quantity": 3
		"promo": "TenPercentOff"
	}
}
*/