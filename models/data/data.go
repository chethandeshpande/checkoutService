package data

type Product struct {
	Name string
	Price float64
}

type AdditionalDetails struct {
	FreeProduct Product
}

type PromotionDetails struct {
	PromotionName      string
	ProductName        string
	MinimumQuantity    int
	DiscountPercentage float64
	DiscountAmount     float64
	AdditionalDetails  AdditionalDetails
}

type ProductPromotionMap struct {
	Promotions map[string]PromotionDetails
}
