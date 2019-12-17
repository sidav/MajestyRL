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

func (ucd *underConstructionData) isCompleted() bool {
	return ucd.currentConstructedAmount >= ucd.maxConstructedAmount
}

//returns current, max and percent 
func (ucd *underConstructionData) getCompletionValues() (int, int, int) { 
	return ucd.currentConstructedAmount, ucd.maxConstructedAmount, 100*ucd.currentConstructedAmount / ucd.maxConstructedAmount
}
