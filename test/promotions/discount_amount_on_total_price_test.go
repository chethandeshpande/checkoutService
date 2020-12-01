package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
	"testing"
)

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
	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 1 || cartAfterPromo.TotalPrice != 90 {
		t.Errorf("Promotion not applied!")
	}
}
