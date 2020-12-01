package promotions

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/promotions"
	"testing"
)

func TestShouldProvideAProductFree(t *testing.T) {
	provideAFreeProduct := promotions.ProvideAFreeProduct{}

	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
		Quantity: 1,
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
		}	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 2 || cartAfterPromo.Products[1].Name != "Raspberry Pi" {
		t.Errorf("Promotion not applied!")
	}
}

func TestShouldUpdateQuantityOfFreeProductIfAlreadyAdded(t *testing.T) {
	provideAFreeProduct := promotions.ProvideAFreeProduct{}

	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
		Quantity: 1,
	},{
		Name:  "Raspberry Pi",
		Price: 30,
		Quantity: 1,
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
		}	}
	cartAfterPromo := provideAFreeProduct.Apply(cart, promo)

	if len(cartAfterPromo.Products) != 2 ||
		cartAfterPromo.Products[1].Name != "Raspberry Pi" ||
		cartAfterPromo.Products[1].Quantity != 2 {
		t.Errorf("Promotion not applied!")
	}
}
