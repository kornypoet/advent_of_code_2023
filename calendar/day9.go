package calendar

import (
	"bufio"
	"strconv"
	"strings"
)

func init() {
	registry["Day9"] = Day9
}

func Day9(input *bufio.Scanner, part int) (result int) {
	for input.Scan() {
		numberStr := strings.Split(input.Text(), " ")
		numbers := make([]int, len(numberStr))
		for i, char := range numberStr {
			numbers[i], _ = strconv.Atoi(char)
		}

		if part == 1 {

			nextSeq := generateDiff(numbers, []int{})
			prediction := 0
			for _, num := range nextSeq {
				prediction += num
			}
			log.Infof("Prediction %d", prediction)
			result += prediction

		} else {
			nextSeq := generatePreDiff(numbers, []int{})
			prediction := 0
			log.Info(nextSeq)
			for i, num := range nextSeq {
				if i == 0 {
					prediction = num
				} else if i % 2 == 0 {					
					prediction += num
				} else {
					prediction -= num
				}
			}
			log.Infof("Prediction %d", prediction)
			result += prediction

		}
	}
	return
}

func generateDiff(seq []int, final []int) []int {
	final = append(final, seq[len(seq) - 1])
	result := []int{}
	for i := 0; i < len(seq) - 1; i++ {
		next := seq[i + 1] - seq[i]
		result = append(result, next)
	}
	if !checkForMatch(result) {
		return generateDiff(result, final)
	}
	return final
}

func generatePreDiff(seq []int, final []int) []int {
	log.Info(seq)
	final = append(final, seq[0])
	result := []int{}
	for i := 0; i < len(seq) - 1; i++ {
		next := seq[i + 1] - seq[i]
		result = append(result, next)
	}
	if !checkForMatch(result) {
		return generatePreDiff(result, final)
	}
	return final
}

func checkForMatch(seq []int) bool {
	for i := 0; i < len(seq); i++ {
		if seq[i] != 0 {
			return false
		}
	}
	return true
}
