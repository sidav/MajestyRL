package main

type resource struct {
	resType   resourceType
	resAmount int
}

type resourceStock struct {
	amount map[resourceType]int
}

func (from *resourceStock) canSubstract(sub *resourceStock) bool {
	for rtype := range from.amount {
		if from.amount[rtype] < sub.amount[rtype] {
			return false
		}
	}
	return true
}

func (from *resourceStock) substract(sub *resourceStock) {
	for rtype := range from.amount {
		from.amount[rtype] -= sub.amount[rtype]
	}
}
