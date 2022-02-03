package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type caesarCipherItem struct {
	str    string
	key    int
	result string
}

var _ = FDescribe("Caesar Cipher", func() {
	Context("When the data is valid", func() {
		It("passes", func() {
			testData := []caesarCipherItem{
				{str: "xyz", key: 2, result: "zab"},
			}

			for _, testItem := range testData {
				result := easy.CaesarCipherEncryptor(testItem.str, testItem.key)
				Expect(result).To(Equal(testItem.result))
			}
		})
	})
})
