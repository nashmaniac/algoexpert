package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type testData struct {
	array          []int
	sequence       []int
	expectedResult bool
}

var _ = Describe("Valid Subsequence", func() {

	var testDataList []testData = []testData{
		{
			array:          []int{1, 1, 1, 1, 1},
			sequence:       []int{1, 1, 1},
			expectedResult: true,
		},
		{
			array:          []int{5, 1, 22, 25, 6, -1, 8, 10},
			sequence:       []int{1, 6, -1, 10},
			expectedResult: true,
		},
	}
	Context("when the valid input is there", func() {
		It("should return true", func() {
			for _, t := range testDataList {
				result := easy.IsValidSubsequence(t.array, t.sequence)
				Expect(result).To(Equal(t.expectedResult))
			}

		})
	})
})
