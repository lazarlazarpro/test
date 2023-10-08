package pack

import "sort"

type Packer struct {
	PackSizes []int
}

func NewPacker(packSizes []int) *Packer {
	// make sure pack sizes are sorted in descending order
	sort.Sort(sort.Reverse(sort.IntSlice(packSizes)))
	return &Packer{PackSizes: packSizes}
}

// roundNumberOfItems rounds the number of items based on the pack sizes
func (p *Packer) roundNumberOfItems(numberOfItems int) int {
	roundedNumberOfItems := numberOfItems
	// take greedy approach
	for _, packSize := range p.PackSizes {
		if numberOfItems >= packSize {
			numberOfItems %= packSize
		}
	}
	// round the total number of items to match the smallest pack size
	if numberOfItems > 0 {
		roundedNumberOfItems += p.PackSizes[len(p.PackSizes)-1] - numberOfItems
	}

	return roundedNumberOfItems
}

// GetPackSizes calculates the pack sizes needed for the given number of items
func (p *Packer) GetPackSizes(numberOfItems int) *map[int]int {
	numberOfItems = p.roundNumberOfItems(numberOfItems)
	packCounts := make(map[int]int)

	// after rounding, the greedy approach will not give a remainder
	for _, packSize := range p.PackSizes {
		if numberOfItems >= packSize {
			packCounts[packSize] = numberOfItems / packSize
			numberOfItems %= packSize
		}
	}

	return &packCounts
}
