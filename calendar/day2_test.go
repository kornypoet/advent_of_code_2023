package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day2", func() {
	It("part 1 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`))
		result := 8
		Expect(calendar.Day2(input, 1)).To(Equal(result))
	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`Game 1: 3 blue, 4 red; 1 red, 2 green, 6 blue; 2 green
Game 2: 1 blue, 2 green; 3 green, 4 blue, 1 red; 1 green, 1 blue
Game 3: 8 green, 6 blue, 20 red; 5 blue, 4 red, 13 green; 5 green, 1 red
Game 4: 1 green, 3 red, 6 blue; 3 green, 6 red; 3 green, 15 blue, 14 red
Game 5: 6 red, 1 blue, 3 green; 2 blue, 1 red, 2 green
`))
		result := 2286
		Expect(calendar.Day2(input, 2)).To(Equal(result))
	})
})
