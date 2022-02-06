package medium_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestMedium(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Medium Suite")
}
