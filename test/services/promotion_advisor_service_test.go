package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
	"testing"
)

func TestShouldReturnApplicablePromotionsForTheProductsInTheCart(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
		Quantity: 1,
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
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 1 {
		t.Errorf("Applicable Promotion not returned! Expected: 1")
	}
}

func TestShouldReturnOnlyOnePromotionWhenSameProductIsSelectedMultipleTimes(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
		Quantity: 2,
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
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 1 {
		t.Errorf("Applicable Promotion not returned! Expected: 1")
	}
}

func TestShouldReturnPromotionsRespectiveToTheProductsSelected(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
		Quantity: 1,
	},{
		Name:  "Alexa Speaker",
		Price: 1,
		Quantity: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	macbookPro := data.PromotionDetails{
		PromotionName:      "RaspberryPi",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}
	alexaSpeaker := data.PromotionDetails{
		PromotionName:      "10PercentOff",
		ProductName:        "Alexa Speaker",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
	}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = macbookPro
	promoToProductMap["Alexa Speaker"] = alexaSpeaker
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 2 ||
		applicablePromos[0].ProductName != "Macbook Pro" ||
		applicablePromos[1].ProductName != "Alexa Speaker" {
		t.Errorf("Applicable Promotion not returned! Expected: 2")
	}
}

func TestShouldReturnNoPromosWhenNoPromosAreAvailable(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
	},{
		Name:  "Macbook Pro",
		Price: 1,
	}}
	cart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: 0,
	}
	promoToProductMap := make(map[string]data.PromotionDetails)
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 0 {
		t.Errorf("Applicable Promotion returned! Expected: 0")
	}
}

func TestShouldReturnPromotionWhenThereAreNoProducts(t *testing.T) {
	cart := dto.ShoppingCart{
		Products:   []dto.Product{},
		TotalPrice: 0,
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
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 0 {
		t.Errorf("Promotion returned! Expected: 0")
	}
}

func TestShouldNotApplyPromotionIfMinimumQuantityIsNotSelected(t *testing.T) {
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
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 0 {
		t.Errorf("Promotion returned! Expected: 0")
	}
}

func TestShouldReturnPromotionIfPromotionIsNotForSameProduct(t *testing.T) {
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
		ProductName:        "Google Home",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	applicablePromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(applicablePromos) != 0 {
		t.Errorf("Promotion returned! Expected: 0")
	}
}
