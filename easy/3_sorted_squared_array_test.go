package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("SortedSquaredArray", func() {
	Context("when the valid array is given", func() {
		It("should return a valid sorted array", func() {
			givenArray := []int{-4, -2, 1, 3, 5, 6}
			resultArray := []int{1, 4, 9, 16, 25, 36}

			arr := easy.SortedSquaredAarray(givenArray)
			for i := 0; i < len(resultArray); i++ {
				Expect(arr[i]).To(Equal(resultArray[i]))
			}
		})
	})
})
