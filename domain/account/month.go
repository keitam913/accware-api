package account

import (
	"math"

	"github.com/keitam913/accware-api/domain/person"
)

type Month struct {
	items []Item
}

func (m Month) Items() []Item {
	is := make([]Item, len(m.items))
	is = append(is, m.items...)
	return is
}

func (m Month) Adjustments(personService person.Service) []Adjustment {
	totals := map[string]int{}
	all := 0
	for _, item := range m.items {
		totals[item.PersonID()] += item.Amount()
		all += item.Amount()
	}
	adjustments := make([]Adjustment, len(m.items))
	for personID, total := range totals {
		ratio, err := personService.PaymentRatio(personID)
		if err != nil {
			panic(err)
		}
		amount := int(math.Round(float64(all)*ratio)) - total
		adjustments = append(adjustments, NewAdjustment(personID, amount))
	}
	return adjustments
}

func (m Month) Totals(personService person.Service) []Total {
	totalMap := map[string]int{}
	for _, item := range m.items {
		totalMap[item.PersonID()] += item.Amount()
	}
	for _, adjustment := range m.Adjustments(personService) {
		totalMap[adjustment.PersonID()] += adjustment.Amount()
	}
	var totals []Total
	for personID, amount := range totalMap {
		totals = append(totals, NewTotal(personID, amount))
	}
	return totals
}
