package services

import (
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"checkoutService/services"
	"testing"
)

func TestShouldReturnPotentialPromotionsForTheProductsInTheCart(t *testing.T) {
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
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		}}

	promoToProductMap := make(map[string]data.PromotionDetails)
	promoToProductMap["Macbook Pro"] = promo
	potentialPromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(potentialPromos) != 1 {
		t.Errorf("Applicable Promotion not returned!")
	}
}

func TestShouldReturnOnlyOnePromotionWhenSameProductIsSelectedMultipleTimes(t *testing.T) {
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
	potentialPromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(potentialPromos) != 1 {
		t.Errorf("Applicable Promotion not returned!")
	}
}

func TestShouldReturnPromotionsRespectiveToTheProductsSelected(t *testing.T) {
	products := []dto.Product{{
		Name:  "Macbook Pro",
		Price: 1,
	},{
		Name:  "Alexa Speaker",
		Price: 1,
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
	potentialPromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(potentialPromos) != 2 ||
		potentialPromos[0].ProductName != "Macbook Pro" ||
		potentialPromos[1].ProductName != "Alexa Speaker" {
		t.Errorf("Applicable Promotion not returned!")
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
	potentialPromos := services.GetApplicablePromos(cart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	if len(potentialPromos) != 0 {
		t.Errorf("Applicable Promotion returned!")
	}
}
