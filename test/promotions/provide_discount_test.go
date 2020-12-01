package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
	"testing"
)

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
