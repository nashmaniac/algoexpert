package easy

import (
	"log"
	"math"
	"sort"
)

func TandemBicycle(redShirtSpeeds []int, blueShirtSpeeds []int, fastest bool) int {
	if fastest {
		sort.Sort(sort.Reverse(sort.IntSlice(redShirtSpeeds)))
	} else {
		sort.Ints(redShirtSpeeds)
	}

	sort.Ints(blueShirtSpeeds)

	countSum := 0
	for i := 0; i < len(redShirtSpeeds); i++ {
		countSum += int(math.Max(float64(redShirtSpeeds[i]), float64(blueShirtSpeeds[i])))
	}

	return countSum

}
