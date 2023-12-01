package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day1", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`1abc2
pqr3stu8vwx
a1b2c3d4e5f
treb7uchet`))
		result := 142
		Expect(calendar.Day1(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`))
		result := 281
		Expect(calendar.Day1(input, 2)).To(Equal(result))
	})
})
