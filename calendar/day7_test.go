package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day7", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`))
		result := 6440
		Expect(calendar.Day7(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483
`))
		result := 5905
		Expect(calendar.Day7(input, 2)).To(Equal(result))
	})
})
