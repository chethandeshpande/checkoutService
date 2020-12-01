package main

import (
	"bufio"
	"checkoutService/app"
	"checkoutService/models/data"
	"checkoutService/models/dto"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("Products Available: \n")
	fmt.Print("1. Google Home: $49.99\n")
	fmt.Print("2. Macbook Pro: $5399.99\n")
	fmt.Print("3. Alexa Speaker: $109.50\n")
	fmt.Print("4. Raspberry Pi: $30.00\n\n")
	fmt.Print("Promotions Available: \n")
	fmt.Print("Each sale of a MacBook Pro comes with a free Raspberry Pi B\n")
	fmt.Print("Buy 3 Google Homes for the price of 2\n")
	fmt.Print("Buying more than 3 Alexa Speakers will have a 10% discount on all Alexa speakers\n\n")
	fmt.Print("Please select products and quantity: \n")
	fmt.Print("Enter in 1:[quantity],2:[quantity],3:[quantity] format to select products and quantity: \n")
	reader := bufio.NewReader(os.Stdin)
	input, err := reader.ReadString('\n')
	if err != nil {
		fmt.Print("Error while processing input\n")
		return
	}
	macbookPro := dto.Product{
		Name:  "Macbook Pro",
		Price: 5399.99,
	}
	alexaSpeaker := dto.Product{
		Name:  "Alexa Speaker",
		Price: 109.50,
	}
	googleHome := dto.Product{
		Name:  "Google Home",
		Price: 49.99,
	}
	raspberryPi := dto.Product{
		Name:  "Raspberry Pi",
		Price: 30.00,
	}
	productsAvailable := map[string]dto.Product{
		"1": googleHome,
		"2": macbookPro,
		"3": alexaSpeaker,
		"4": raspberryPi,
	}
	input = strings.ReplaceAll(input, "\n", "")
	productsSelected := strings.Split(input, ",")
	var products []dto.Product
	var totalPrice float64
	for _, product := range productsSelected {
		productDetails := strings.Split(product, ":")
		product := productsAvailable[productDetails[0]]
		quantity, err := strconv.Atoi(productDetails[1])
		if err != nil {
			fmt.Print("Error while processing input\n")
			return
		}
		product.Quantity = quantity
		products = append(products, product)
		totalPrice += product.Price * math.Round(float64(product.Quantity)*100)/100
	}

	shoppingCart := dto.ShoppingCart{
		Products:   products,
		TotalPrice: totalPrice,
	}
	fmt.Print("Shopping Before After applying Promotions: \n")
	printCart(shoppingCart)
	macbookProPromo := data.PromotionDetails{
		PromotionName:      "RaspberryPiFree",
		ProductName:        "Macbook Pro",
		MinimumQuantity:    1,
		DiscountPercentage: 0,
		AdditionalDetails: data.AdditionalDetails{
			FreeProduct: data.Product{
				Name:  "Raspberry Pi",
				Price: 30,
			},
		},
	}
	alexaSpeakerPromo := data.PromotionDetails{
		PromotionName:      "TenPercentOff",
		ProductName:        "Alexa Speaker",
		MinimumQuantity:    3,
		DiscountPercentage: 10,
	}
	googleHomePromo := data.PromotionDetails{
		PromotionName:      "BuyTwoGetOne",
		ProductName:        "Google Home",
		MinimumQuantity:    2,
		DiscountAmount:     googleHome.Price,
		DiscountPercentage: 0,
	}

	promoToProductMap := map[string]data.PromotionDetails{
		"Macbook Pro":   macbookProPromo,
		"Alexa Speaker": alexaSpeakerPromo,
		"Google Home":   googleHomePromo,
	}

	cartAfterAppliedPromos := app.Checkout(shoppingCart, data.ProductPromotionMap{
		Promotions: promoToProductMap,
	})

	fmt.Print("Shopping Cart After applying Promotions: \n")
	printCart(cartAfterAppliedPromos)
	printPromotionsApplied(cartAfterAppliedPromos)
}

func printPromotionsApplied(cartAfterAppliedPromos dto.ShoppingCart) {
	fmt.Print("Applied Promotions: \n")
	for _, promo := range cartAfterAppliedPromos.AdditionalDetails.AppliedPromos {
		fmt.Printf("%s on %s\n", promo.PromotionName, promo.ProductName)
	}
}

func printCart(cartAfterAppliedPromos dto.ShoppingCart) {
	fmt.Print("Product, Price, Quantity\n")
	for _, product := range cartAfterAppliedPromos.Products {
		fmt.Printf("%s, %.2f, %d\n", product.Name, math.Round(product.Price*1000)/1000, product.Quantity)
	}
	fmt.Print("Total Price: ", cartAfterAppliedPromos.TotalPrice, "\n\n")
}
