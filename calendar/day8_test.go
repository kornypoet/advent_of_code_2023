package calendar_test

import (
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	"advent_of_code/calendar"
	"bufio"
	"strings"
)

var _ = Describe("Day8", func() {
//	It("part 1 should be correct", func() {
//		input := bufio.NewScanner(strings.NewReader(`AAA = (BBB, CCC)
// BBB = (DDD, EEE)
// CCC = (ZZZ, GGG)
// DDD = (DDD, DDD)
// EEE = (EEE, EEE)
// GGG = (GGG, GGG)
// ZZZ = (ZZZ, ZZZ)
// `))
	//	result := 2
	//	Expect(calendar.Day8(input, 1)).To(Equal(result))
	// })

//	It("part 1 should be correct, second example", func() {
//		input := bufio.NewScanner(strings.NewReader(`AAA = (BBB, BBB)
// BBB = (AAA, ZZZ)
// ZZZ = (ZZZ, ZZZ)
// `))
//		result := 6
//		Expect(calendar.Day8(input, 1)).To(Equal(result))
//	})

	It("part 2 should be correct", func() {
		input := bufio.NewScanner(strings.NewReader(`11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)
`))
		result := 6
		Expect(calendar.Day8(input, 2)).To(Equal(result))
	})
})
