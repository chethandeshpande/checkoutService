package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
	"testing"
)

func TestShouldApplyApplicablePromotions(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
	},{
		Name:  "Macbook Pro",
		Price: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 1,
	}
	promo := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}	}
	cartAfterPromo := services.ApplyPromos(cart, []data.PromotionDetails{promo})

	if len(cartAfterPromo.Products) != 3 || cartAfterPromo.TotalPrice != cart.TotalPrice {
		t.Errorf("Applicable promotion not applied!")
	}
}