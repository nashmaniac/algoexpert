package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type tandemBicycleItem struct {
	blueShirtSpeeds []int
	fastest         bool
	redShirtSpeeds  []int
	result          int
}

var _ = Describe("tandem bicycle", func() {
	Context("when the data is correct", func() {
		It("passes", func() {
			testData := []tandemBicycleItem{
				{
					blueShirtSpeeds: []int{3, 6, 7, 2, 1},
					redShirtSpeeds:  []int{5, 5, 3, 9, 2},
					fastest:         false,
					result:          25,
				},
				{
					blueShirtSpeeds: []int{3, 3, 4, 6, 1, 2},
					fastest:         false,
					redShirtSpeeds:  []int{1, 2, 1, 9, 12, 3},
					result:          30,
				},
			}

			for _, testItem := range testData {
				r := easy.TandemBicycle(testItem.redShirtSpeeds, testItem.blueShirtSpeeds, testItem.fastest)
				Expect(r).To(Equal(testItem.result))
			}
		})
	})
})
