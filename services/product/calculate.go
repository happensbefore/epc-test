package product

import (
	"sort"
	"strconv"

	. "epc-test/models"
)

func isConditionContainsRule(rule RuleApplicability, conditions []Condition) bool {
	for _, condition := range conditions {
		if condition.RuleName == rule.CodeName && condition.Value == rule.Value {
			return true
		}
	}

	return false
}

func isPriceApplicableByRule(price Price, conditions []Condition) bool {
	for _, rule := range price.RuleApplicabilities {
		if !isConditionContainsRule(rule, conditions) {
			return false
		}
	}

	return true
}

func isThereHaveMainComponent(components []Component) bool {
	if components != nil && len(components) > 0 {
		for _, c := range components {
			if c.IsMain {
				return true
			}
		}
	}

	return false
}

func getComponentApplicableCostPrices(prices []Price, conditions []Condition) []Price {
	var ap []Price

	for _, price := range prices {
		if price.PriceType == PriceTypeCost && isPriceApplicableByRule(price, conditions) {
			ap = append(ap, price)
		}
	}

	return ap
}

func isDiscountApplicableToPrice(discountPrice Price, costPrice Price) bool {
	for _, discountRule := range discountPrice.RuleApplicabilities {
		for _, costRule := range costPrice.RuleApplicabilities {
			if discountRule.CodeName == costRule.CodeName {
				dv, _ := strconv.Atoi(discountRule.Value)
				cv, _ := strconv.Atoi(costRule.Value)

				switch discountRule.Operator {
				case OperatorGreaterThanOrEqual:
					return cv >= dv
				case OperatorLessThanOrEqual:
					return cv <= dv
				default:
					return dv == cv
				}
			}
		}
	}

	return false
}
func getDiscountForPrice(discountPrices []Price, costPrice Price) float64 {
	sort.Slice(discountPrices, func(i, j int) bool {
		return discountPrices[i].Cost > discountPrices[j].Cost
	})

	for _, discountPrice := range discountPrices {
		if isDiscountApplicableToPrice(discountPrice, costPrice) {
			return discountPrice.Cost
		}
	}

	return 0
}

func setComponentDiscount(component *Component, costPrices []Price) {
	if len(costPrices) == 0 {
		return
	}

	var discountPrices []Price

	for _, price := range component.Prices {

		if price.PriceType == PriceTypeDiscount {
			discountPrices = append(discountPrices, price)
		}
	}

	if len(discountPrices) == 0 {
		return
	}

	for icp, costPrice := range costPrices {
		discountPrice := getDiscountForPrice(discountPrices, costPrice)

		if discountPrice > 0 {
			costPrices[icp].Cost = costPrice.Cost - costPrice.Cost*discountPrice/100
		}
	}
}

func computeTotalCost(components []Component) Price {
	var totalCost Price

	for _, c := range components {
		for _, p := range c.Prices {
			totalCost.Cost += p.Cost
		}
	}

	return totalCost
}

func Calculate(product *Product, conditions []Condition) (offer *Offer, err error) {
	if product == nil || conditions == nil {
		return nil, nil
	}

	if !isThereHaveMainComponent(product.Components) || len(conditions) == 0 {
		return &Offer{}, nil
	}

	var components []Component

	for _, component := range product.Components {
		componentPrices := getComponentApplicableCostPrices(component.Prices, conditions)

		setComponentDiscount(&component, componentPrices)

		if len(componentPrices) > 0 {
			components = append(components, Component{Name: component.Name, Prices: componentPrices, IsMain: component.IsMain})
		}
	}

	if !isThereHaveMainComponent(components) {
		return nil, nil
	}

	offer = &Offer{
		Product: Product{
			Name:       product.Name,
			Components: components,
		},
		TotalCost: computeTotalCost(components),
	}

	return offer, nil
}
