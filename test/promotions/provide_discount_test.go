package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
	"testing"
)

func TestShouldNotProvideDiscountIfThereAreNoProducts(t *testing.T) {
	provideAFreeProduct := promotions.ProvideDiscount{}

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
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 0 && cartAfterPromo.TotalPrice == cart.TotalPrice{
		t.Errorf("Promotion applied when no products are present!")
	}
}

func TestShouldNotProvideDiscountIfPromotionIsNotForSameProduct(t *testing.T) {
	provideAFreeProduct := promotions.ProvideDiscount{}

	products := []dto.Product{{
		Name:  "Alexa Speakers",
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
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 && cartAfterPromo.TotalPrice == cart.TotalPrice{
		t.Errorf("Promotion applied when promotion was for different product!")
	}
}

func TestShouldNotProvideDiscountIfMinimumQuantityIsNotMet(t *testing.T) {
	provideAFreeProduct := promotions.ProvideDiscount{}

	products := []dto.Product{{
		Name:  "Alexa Speakers",
		Price: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    2,
		DiscountPercentage: 0,
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 && cartAfterPromo.TotalPrice == cart.TotalPrice{
		t.Errorf("Promotion applied when minimum quantity is not met!")
	}
}

func TestShouldProvideDiscount(t *testing.T) {
	provideAFreeProduct := promotions.ProvideDiscount{}

	products := []dto.Product{{
		Name:  "Alexa Speakers",
		Price: 100,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 100,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 || cartAfterPromo.TotalPrice != 90 {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldProvideDiscountWhenMoreThanMinimumQuantityIsSelected(t *testing.T) {
	provideAFreeProduct := promotions.ProvideDiscount{}

	products := []dto.Product{{
		Name:  "Alexa Speakers",
		Price: 100,
	},{
		Name:  "Alexa Speakers",
		Price: 100,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 200,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Alexa Speakers",
		MinimumQuantity:    1,
		DiscountPercentage: 10,
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 2 || cartAfterPromo.TotalPrice != 180 {
		t.Errorf("Promotion not applied!")
	}
}
