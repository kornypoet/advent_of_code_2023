package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day3", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`))
		result := 4361
		Expect(calendar.Day3(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`467..114..
...*......
..35..633.
......#...
617*......
.....+.58.
..592.....
......755.
...$.*....
.664.598..
`))
		result := 467835
		Expect(calendar.Day3(input, 2)).To(Equal(result))
	})
})
