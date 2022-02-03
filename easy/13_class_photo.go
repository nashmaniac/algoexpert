package easy

import (
	"sort"
)

func ClassPhotos(redShirtHeights []int, blueShirtHeights []int) bool {
	sort.Ints(redShirtHeights)
	sort.Ints(blueShirtHeights)

	isBlueFront := blueShirtHeights[0] <= redShirtHeights[0]

	for i := 0; i < len(redShirtHeights); i++ {
		if isBlueFront {
			if redShirtHeights[i] < blueShirtHeights[i] {
				return false
			}	
		} else {
			if redShirtHeights[i] > blueShirtHeights[i] {
				return false
			}	
		}
	}
	return true
}
