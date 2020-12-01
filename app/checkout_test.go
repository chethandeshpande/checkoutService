package app

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"math"
	"testing"
)

func TestShouldProvideAFreeRaspberryPiOnBuyingAMacbookPro(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 5399.99,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaDiscountPromotion := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
	}
	googleHomePromotion := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     10,
	}

	promotions := map[string]data.PromotionDetails{
		"Macbook Pro":    macbookProPromotion,
		"Alexa Speakers": alexaDiscountPromotion,
		"Google Home":    googleHomePromotion,
	}
	cartAfterPromo := Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if len(cartAfterPromo.Products) != 2 ||
		cartAfterPromo.TotalPrice != cart.TotalPrice ||
		cartAfterPromo.Products[1].Name != "Raspberry Pi" {
		t.Errorf("Macbook Pro Promotion not Applied!")
	}
}

func TestShouldProvideTenPercentDiscountOnBuyingThreeAlexaSpeakers(t *testing.T) {
	products := []dto.Product{{
		Name:  "Alexa Speakers",
		Price: 109.50,
	},{
		Name:  "Alexa Speakers",
		Price: 109.50,
	},{
		Name:  "Alexa Speakers",
		Price: 109.50,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 328.5,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaDiscountPromotion := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    2,
		DiscountPercentage: 10,
	}
	googleHomePromotion := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     10,
	}

	promotions := map[string]data.PromotionDetails{
		"Macbook Pro":    macbookProPromotion,
		"Alexa Speakers": alexaDiscountPromotion,
		"Google Home":    googleHomePromotion,
	}
	cartAfterPromo := Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 295.65); diff > 0.000001 || len(cartAfterPromo.Products) != 3 {
		t.Errorf("Alexa Speakers Promotion not Applied!")
	}
}

func TestShouldProvideThreeGoogleHomesAtThePriceOfTwo(t *testing.T) {
	products := []dto.Product{{
		Name:  "Google Home",
		Price: 49.99,
	},{
		Name:  "Google Home",
		Price: 49.99,
	},{
		Name:  "Alexa Speakers",
		Price: 49.99,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 149.97,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaDiscountPromotion := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    2,
		DiscountPercentage: 10,
	}
	googleHomePromotion := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     49.99,
	}

	promotions := map[string]data.PromotionDetails{
		"Macbook Pro":    macbookProPromotion,
		"Alexa Speakers": alexaDiscountPromotion,
		"Google Home":    googleHomePromotion,
	}
	cartAfterPromo := Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 99.98); diff > 0.000001 || len(cartAfterPromo.Products) != 3 {
		t.Errorf("Google Home Promotion not Applied!")
	}
}

func TestShouldNotApplyAnyPromotion(t *testing.T) {
	products := []dto.Product{{
		Name:  "Raspberry Pi",
		Price: 30.00,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 30.00,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaDiscountPromotion := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    2,
		DiscountPercentage: 10,
	}
	googleHomePromotion := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     49.99,
	}

	promotions := map[string]data.PromotionDetails{
		"Macbook Pro":    macbookProPromotion,
		"Alexa Speakers": alexaDiscountPromotion,
		"Google Home":    googleHomePromotion,
	}
	cartAfterPromo := Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 30.00); diff > 0.000001 || len(cartAfterPromo.Products) != 1 {
		t.Errorf("Promotion Applied!")
	}
}

func TestShouldApplyMultiplePromotions(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 5399.99,
	},{
		Name:  "Google Home",
		Price: 49.99,
	},{
		Name:  "Google Home",
		Price: 49.99,
	},{
		Name:  "Alexa Speakers",
		Price: 49.99,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 5549.96,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	alexaDiscountPromotion := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    2,
		DiscountPercentage: 10,
	}
	googleHomePromotion := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     49.99,
	}

	promotions := map[string]data.PromotionDetails{
		"Macbook Pro":    macbookProPromotion,
		"Alexa Speakers": alexaDiscountPromotion,
		"Google Home":    googleHomePromotion,
	}
	cartAfterPromo := Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 5499.97); diff > 0.000001 || len(cartAfterPromo.Products) != 5 {
		t.Errorf("Macbook Pro and Google Home promotions not applied!")
	}
}
