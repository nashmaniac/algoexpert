package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type nonConstructibleChangeItem struct {
	coins  []int
	result int
}

var _ = Describe("Non Destructible Change", func() {
	Context("When the correct input is given", func() {
		It("passes", func() {
			testData := []nonConstructibleChangeItem{
				{coins: []int{5, 7, 1, 1, 2, 3, 22}, result: 20},
				{coins: []int{1, 1, 4}, result: 3},
			}

			for i := 0; i < len(testData); i++ {
				result := easy.NonConstructibleChange(testData[i].coins)
				Expect(result).To(Equal(testData[i].result))
			}
		})
	})
})
