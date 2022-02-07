package medium_test

import (
	"log"

	"github.com/nashmaniac/algoexpert/medium"
	. "github.com/onsi/ginkgo"
	// . "github.com/onsi/gomega"
)

var _ = FDescribe("Merge Interval", func() {
	Context("when the data is valid", func() {
		It("passes", func() {
			testData := [][]int{
				{1, 2}, {3, 5}, {4, 7}, {6, 8}, {9, 10},
			}

			l := medium.MergeOverlappingIntervals(testData)
			log.Println(l)
		})
	})
})
