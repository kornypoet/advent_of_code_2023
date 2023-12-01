package calendar

import (
	"bufio"
	"regexp"
	"strconv"
)

func init() {
	registry["Day1"] = Day1
}

func convert(num string) string {
	lookup := map[string]string{
		"one":   "1",
		"two":   "2",
		"three": "3",
		"four":  "4",
		"five":  "5",
		"six":   "6",
		"seven": "7",
		"eight": "8",
		"nine":  "9",
	}
	replace, ok := lookup[num]
	if ok {
		return replace
	} else {
		return num
	}
}

func Day1(input *bufio.Scanner, part int) (result int) {
	var re *regexp.Regexp
	if part == 1 {
		re = regexp.MustCompile(`\d`)
	} else {
		re = regexp.MustCompile(`(\d|one|two|three|four|five|six|seven|eight|nine)`)
	}

	for input.Scan() {
		match := re.FindAll([]byte(input.Text()), -1)
		if match != nil {
			first := string(match[0])
			last := string(match[len(match)-1])
			number, _ := strconv.Atoi(convert(first) + convert(last))
			log.Debug(number)
			result += number
		}
	}
	return
}
