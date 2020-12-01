package app

import (
	"checkoutService/app"
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"math"
	"testing"
)

func TestShouldProvideAFreeRaspberryPiOnBuyingAMacbookPro(t *testing.T) {
	products := []dto.Product{{
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
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
	cartAfterPromo := app.Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if len(cartAfterPromo.Products) != 2 ||
		cartAfterPromo.TotalPrice != cart.TotalPrice ||
		cartAfterPromo.Products[1].Name != "Raspberry Pi" ||
		len(cartAfterPromo.AdditionalDetails.AppliedPromos) != 1 ||
		cartAfterPromo.AdditionalDetails.AppliedPromos[0].PromotionName != "RaspberryPiFree" {
		t.Errorf("Macbook Pro Promotion not Applied!")
	}
}

func TestShouldProvideTenPercentDiscountOnBuyingThreeAlexaSpeakers(t *testing.T) {
	products := []dto.Product{{
		Name:     "Alexa Speakers",
		Price:    109.50,
		Quantity: 3,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 328.5,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
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
	cartAfterPromo := app.Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 295.65); diff > 0.000001 ||
		len(cartAfterPromo.Products) != 1 ||
		len(cartAfterPromo.AdditionalDetails.AppliedPromos) != 1 {
		t.Errorf("Alexa Speakers Promotion not Applied!")
	}
}

func TestShouldProvideThreeGoogleHomesAtThePriceOfTwo(t *testing.T) {
	products := []dto.Product{{
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 3,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 149.97,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
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
	cartAfterPromo := app.Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 99.98); diff > 0.000001 ||
		len(cartAfterPromo.Products) != 1 ||
		len(cartAfterPromo.AdditionalDetails.AppliedPromos) != 1 {
		t.Errorf("Google Home Promotion not Applied!")
	}
}

func TestShouldNotApplyAnyPromotion(t *testing.T) {
	products := []dto.Product{{
		Name:     "Raspberry Pi",
		Price:    30.00,
		Quantity: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 30.00,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
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
	cartAfterPromo := app.Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 30.00); diff > 0.000001 ||
		len(cartAfterPromo.Products) != 1 ||
		len(cartAfterPromo.AdditionalDetails.AppliedPromos) != 0 {
		t.Errorf("Promotion Applied!")
	}
}

func TestShouldApplyMultiplePromotions(t *testing.T) {
	products := []dto.Product{{
		Name:     "Macbook Pro",
		Price:    5399.99,
		Quantity: 1,
	}, {
		Name:     "Google Home",
		Price:    49.99,
		Quantity: 2,
	}, {
		Name:     "Alexa Speakers",
		Price:    49.99,
		Quantity: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 5549.96,
	}
	macbookProPromotion := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
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
	cartAfterPromo := app.Checkout(cart, data.ProductPromotionMap{Promotions: promotions})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 5499.97); diff > 0.000001 ||
		len(cartAfterPromo.Products) != 4 ||
		len(cartAfterPromo.AdditionalDetails.AppliedPromos) != 2 {
		t.Errorf("Macbook Pro and Google Home promotions not applied!")
	}
}
