package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"testing"
)

func TestShouldProvideAProductFree(t *testing.T) {
	provideAFreeProduct := ProvideAFreeProduct{}

	products := []dto.Product{{
		Name:  "Macbook Pro",
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

	if len(cartAfterPromo.Products) != 2 || cartAfterPromo.Products[1].Name != "Raspberry Pi" {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldProvideAProductFreeWhenMoreThanMinimumQuantityIsProvided(t *testing.T) {
	provideAFreeProduct := ProvideAFreeProduct{}

	products := []dto.Product{
		{
			Name:  "Macbook Pro",
			Price: 1,
		},
		{
			Name:  "Macbook Pro",
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

	if len(cartAfterPromo.Products) != 3 || cartAfterPromo.Products[2].Name != "Raspberry Pi" {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldNotApplyPromotionWhenThereAreNoProducts(t *testing.T) {
	provideAFreeProduct := ProvideAFreeProduct{}

	cart := dto.ShoppingCart{
		Products:   []dto.Product{},
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

	if len(cartAfterPromo.Products) != 0 {
		t.Errorf("Promotion applied without product!")
	}
}

func TestShouldNotApplyPromotionIfMinimumQuantityIsNotSelected(t *testing.T) {
	provideAFreeProduct := ProvideAFreeProduct{}
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
	}}

	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    2,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != len(cart.Products) {
		t.Errorf("Promotion applied when minimum quantity is not met!")
	}
}

func TestShouldNotApplyPromotionIfPromotionIsNotForSameProduct(t *testing.T) {
	provideAFreeProduct := ProvideAFreeProduct{}
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
	}}

	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Google home",
		MinimumQuantity:    2,
		DiscountPercentage: 0,
		AdditionalDetails: map[string]string{
			"FreeProduct": "Raspberry Pi",
		},
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != len(cart.Products) {
		t.Errorf("Promotion applied when Product is different")
	}
}
