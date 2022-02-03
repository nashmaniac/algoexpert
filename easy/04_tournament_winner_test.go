package easy_test

import (
	"github.com/nashmaniac/algoexpert/easy"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

type tournamentWinnerTestItem struct {
	competitions [][]string
	result       []int
	winner       string
}

var _ = Describe("Tournament Winner", func() {
	Context("when the valid input is given", func() {
		It("returns the correct winner", func() {
			testData := []tournamentWinnerTestItem{
				{
					competitions: [][]string{
						{"HTML", "C#"},
						{"C#", "Python"},
						{"Python", "HTML"},
					},
					result: []int{0, 0, 1},
					winner: "Python",
				},
			}
			for i := 0; i < len(testData); i++ {
				w := easy.TournamentWinner(testData[i].competitions, testData[i].result)
				Expect(w).To(Equal(testData[i].winner))
			}
		})
	})
})
