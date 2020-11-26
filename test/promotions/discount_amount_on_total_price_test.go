package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
	"testing"
)

func TestShouldNotWaiveOneProductPriceIfThereAreNoProducts(t *testing.T) {
	provideAFreeProduct := promotions.DiscountAmountOnTotalPrice{}

	var products []dto.Product
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 0 && cartAfterPromo.TotalPrice == cart.TotalPrice {
		t.Errorf("Promotion applied when no products are present!")
	}
}

func TestShouldNotWaiveOneProductPriceIfPromotionIsNotForSameProduct(t *testing.T) {
	provideAFreeProduct := promotions.DiscountAmountOnTotalPrice{}

	products := []dto.Product{{
		Name:  "Google Home",
		Price: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 && cartAfterPromo.TotalPrice == cart.TotalPrice {
		t.Errorf("Promotion applied when promotion was for different product!")
	}
}

func TestShouldNotWaiveOneProductPriceIfMinimumQuantityIsNotMet(t *testing.T) {
	provideAFreeProduct := promotions.DiscountAmountOnTotalPrice{}

	products := []dto.Product{{
		Name:  "Google Home",
		Price: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Google Home",
		MinimumQuantity:    2,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 && cartAfterPromo.TotalPrice == cart.TotalPrice {
		t.Errorf("Promotion applied when minimum quantity is not met!")
	}
}

func TestShouldWaiveOneProductPrice(t *testing.T) {
	provideAFreeProduct := promotions.DiscountAmountOnTotalPrice{}

	products := []dto.Product{{
		Name:  "Google Home",
		Price: 100,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 100,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     10,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 || cartAfterPromo.TotalPrice != 90 {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldWaiveOneProductPriceWhenMoreThanMinimumQuantityIsSelected(t *testing.T) {
	provideAFreeProduct := promotions.DiscountAmountOnTotalPrice{}

	products := []dto.Product{{
		Name:  "Google Home",
		Price: 100,
	}, {
		Name:  "Google Home",
		Price: 100,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 200,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
		DiscountAmount:     10,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 2 || cartAfterPromo.TotalPrice != 190 {
		t.Errorf("Promotion not applied!")
	}
}
