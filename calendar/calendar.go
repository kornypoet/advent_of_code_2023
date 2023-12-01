package calendar

import (
	"bufio"
	"fmt"
	"os"

	"github.com/sirupsen/logrus"
)

var log = logrus.New()
var registry = make(map[string]func(*bufio.Scanner, int) int)

var input = `two1nine
eightwothree
abcone2threexyz
xtwone3four
4nineeightseven2
zoneight234
7pqrstsixteen`

func inputFile(day string) *bufio.Scanner {
	file, err := os.Open(fmt.Sprintf("input/%s.txt", day))
	if err != nil {
		log.Fatal(err)
	}
	return bufio.NewScanner(file)
}

func Run(day string, part int, debug bool) {
	if debug {
		log.SetLevel(logrus.DebugLevel)
	}
	if day == "0" {
		for selection, fn := range registry {
			log.Infof("Running %s", selection)
			for _, p := range []int{1, 2} {
				result := fn(inputFile(selection), p)
				log.Infof("Part %d result: %d", p, result)
			}
		}
	} else {
		selection := "Day" + day
		log.Infof("Running %s", selection)
		result := registry[selection](inputFile(selection), part)
		log.Infof("Part %d result: %d", part, result)
	}
}
