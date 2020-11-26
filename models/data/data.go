package data

type PromotionDetails struct {
	PromotionName      string
	ProductName        string
	MinimumQuantity    int
	DiscountPercentage float64
	DiscountAmount     float64
	AdditionalDetails  map[string]string
}

type ProductPromotionMap struct {
	Promotions map[string]PromotionDetails
}
