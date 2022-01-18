package algoexpert_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestAlgoexpert(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Algoexpert Suite")
}
