package medium_test

import (
	"log"

	"github.com/nashmaniac/algoexpert/medium"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = FDescribe("Spiral Traverse", func() {
	Context("when the data is valid", func() {
		It("works", func() {

			array := [][]int{
				{4, 2, 3, 6, 7, 8, 1, 9, 5, 10},
				{12, 19, 15, 16, 20, 18, 13, 17, 11, 14},
			}
			arr := medium.SpiralTraverse(array)
			log.Println(arr)
		})
	})
})
