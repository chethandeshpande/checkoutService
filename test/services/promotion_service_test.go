package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
	"math"
	"testing"
)

func TestShouldApplyApplicablePromotions(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 100,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 100,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}
	cartAfterPromo := services.ApplyPromos(cart, []data.PromotionDetails{promo})

	if diff := math.Abs(cartAfterPromo.TotalPrice - 100); diff > 0.000001 || len(cartAfterPromo.Products) != 2 {
		t.Errorf("Applicable promotion not applied! Expected: RaspberryPi to be applied")
	}
}
