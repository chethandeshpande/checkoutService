package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
	"testing"
)

func TestShouldReturnPromotionStrategy(t *testing.T) {
	strategy := services.GetPromotionStrategy("BuyTwoGetOne")

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
	}

	cartAfterPromo := strategy.Apply(cart, promo)
	if len(cartAfterPromo.Products) != 1 || cartAfterPromo.TotalPrice != 90 {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldReturnDefaultPromotionStrategy(t *testing.T) {
	strategy := services.GetPromotionStrategy("SomePromotion")

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
	}

	cartAfterPromo := strategy.Apply(cart, promo)
	if len(cartAfterPromo.Products) != 1 || cartAfterPromo.TotalPrice != 100 {
		t.Errorf("Promotion not applied!")
	}
}
