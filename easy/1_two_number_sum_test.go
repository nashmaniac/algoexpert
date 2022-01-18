package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Two number sum", func() {

	var arrayList []int
	var targetSum int

	Context("when the array does not contain the sum", func() {
		It("should return nil", func() {
			arrayList = []int{3, 5, -4, 8, 11, 1, -1, 6}
			targetSum = 10

			result := easy.TwoNumberSum(arrayList, targetSum)
			Ω(result).Should(ContainElement(-1))
			Ω(result).Should(ContainElement(11))
		})
	})
})

var _ = Describe("Two number sum sorting method", func() {
	var arrayList []int
	var targetSum int

	Context("when the array does not contain the sum", func() {
		It("should return nil", func() {
			arrayList = []int{3, 5, -4, 8, 11, 1, -1, 6}
			targetSum = 10

			result := easy.TwoNumberSumSortingMethod(arrayList, targetSum)
			Ω(result).Should(ContainElement(-1))
			Ω(result).Should(ContainElement(11))
		})
	})
})
