package main

type underConstructionData struct {
	currentConstructedAmount, maxConstructedAmount int
	isSelfConstructing                             bool
}

func (ucd *underConstructionData) clone() *underConstructionData {
	newUcd := underConstructionData{
		currentConstructedAmount: ucd.currentConstructedAmount,
		maxConstructedAmount: ucd.maxConstructedAmount,
		isSelfConstructing: ucd.isSelfConstructing,
	}
	return &newUcd
}
