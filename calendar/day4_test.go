package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day4", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(``))
		result := 0
		Expect(calendar.Day4(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(``))
		result := 0
		Expect(calendar.Day4(input, 2)).To(Equal(result))
	})
})
