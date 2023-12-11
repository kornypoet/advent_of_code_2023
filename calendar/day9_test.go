package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day9", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`))
		result := 114
		Expect(calendar.Day9(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45
`))

		result := 2
		Expect(calendar.Day9(input, 2)).To(Equal(result))
	})
})
