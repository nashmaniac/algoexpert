package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type runLengthEncodingItem struct {
	str    string
	result string
}

var _ = Describe("RunLengthEncoding", func() {
	Context("when the data is valid", func() {
		It("passes", func() {

			testData := []runLengthEncodingItem{
				{str: "AAAAAAAAAAAAABBCCCCDD", result: "9A4A2B4C2D"},
			}

			for _, testItem := range testData {
				r := easy.RunLengthEncoding(testItem.str)
				Expect(r).To(Equal(testItem.result))
			}

		})
	})
})
