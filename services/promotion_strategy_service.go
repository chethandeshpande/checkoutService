package services

import (
	"checkoutService/promotions"
	"sync"
)

type promotionStrategyMap map[string]promotions.Promotion

var once sync.Once

var (
	strategyMap promotionStrategyMap
)

func getAllPromotionStrategies() promotionStrategyMap {
	once.Do(func() {
		strategyMap = map[string]promotions.Promotion{
			"BuyTwoGetOne":    promotions.DiscountAmountOnTotalPrice{},
			"RaspberryPiFree": promotions.ProvideAFreeProduct{},
			"TenPercentOff":   promotions.ProvideDiscount{},
		}
	})

	return strategyMap
}

func GetPromotionStrategy(promoName string) promotions.Promotion {
	strategies := getAllPromotionStrategies()
	if strategies[promoName] != nil {
		return strategies[promoName]
	}
	return promotions.DefaultPromotion{}
}
