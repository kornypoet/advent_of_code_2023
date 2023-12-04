package calendar

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry["Day4"] = Day4
}

type Card struct {
	id   int
	wins int
}

func Day4(input *bufio.Scanner, part int) (result int) {
	totalCards := []*Card{}
	for input.Scan() {
		info := strings.Split(input.Text(), ":")
		re := regexp.MustCompile(`^Card\s+(\d+)`)
		cardInfo := re.FindStringSubmatch(info[0])
		cardId, _ := strconv.Atoi(cardInfo[1])
		card := &Card{ id: cardId, wins: 0 }
		numbers := strings.Split(info[1], "|")
		picks := strings.Fields(numbers[0])
		winners := strings.Fields(numbers[1])
		points := 0
		for _, pick := range picks {
			for _, winner := range winners {
				if pick == winner {
					card.wins++
					if points == 0 {
						points = 1
					} else {
						points = points * 2
					}
				}
			}
		}
		result += points
		totalCards = append(totalCards, card)
	}
	if part == 1 {
		return
	} else {
		result = 0
		process := map[int]int{}
		for _, card := range totalCards {
			process[card.id] = 1
		}
		for _, card := range totalCards {
			log.Infof("Processing card %d with %d wins %d times", card.id, card.wins, process[card.id])
			for c := 0; c < process[card.id]; c++ {
				for i := 0; i < card.wins; i++ {
					process[card.id + i + 1]++
				}
			}
		}
		for _, proc := range process {
			result += proc
		}
		return
	}
}
