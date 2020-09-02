package main

//type resource struct {
//	resType   resourceType
//	resAmount int
//}

type resourceStock struct {
	amount map[resourceType]int
}

func (from *resourceStock) canSubstract(sub *resourceStock) bool {
	for rtype := range sub.amount {
		if amt, exists := from.amount[rtype]; !exists || amt < sub.amount[rtype] {
			return false
		}
	}
	return true
}

func (from *resourceStock) addStock(add *resourceStock) {
	for rtype := range from.amount {
		from.amount[rtype] += add.amount[rtype]
	}
}

func (from *resourceStock) substract(sub *resourceStock) {
	for rtype := range from.amount {
		from.amount[rtype] -= sub.amount[rtype]
	}
}
