package calendar

import (
	"bufio"
	"reflect"
	"sort"
	"strconv"
	"strings"
)

func init() {
	registry["Day7"] = Day7
}

var ranks = map[string]int{
	"five-of-a-kind":  7,
	"four-of-a-kind":  6,
	"full-house":      5,
	"three-of-a-kind": 4,
	"two-pair":        3,
	"one-pair":        2,
	"high-card":       1,
}

var cardRanks = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 10,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

var jokerRanks = map[string]int{
	"A": 13,
	"K": 12,
	"Q": 11,
	"J": 0,
	"T": 9,
	"9": 8,
	"8": 7,
	"7": 6,
	"6": 5,
	"5": 4,
	"4": 3,
	"3": 2,
	"2": 1,
}

func upgrade(label string, jokers int) int {
	switch label {
	case "five-of-a-kind":
		return ranks[label]
	case "four-of-a-kind":
		if jokers == 1 || jokers == 4 {
			return ranks["five-of-a-kind"]
		} else {
			return ranks[label]
		}
	case "full-house":
		if jokers == 2 || jokers == 3 {
			return ranks["five-of-a-kind"]
		} else {
			return ranks[label]
		}
	case "three-of-a-kind":
		if jokers == 1 || jokers == 3 {
			return ranks["four-of-a-kind"]
		} else {
			return ranks[label]
		}
	case "two-pair":
		if jokers == 1 {
			return ranks["full-house"]
		} else if jokers == 2 {
			return ranks["four-of-a-kind"]
		} else {
			return ranks[label]
		}
	case "one-pair":
		if jokers == 1 || jokers == 2 {
			return ranks["three-of-a-kind"]
		} else {
			return ranks[label]
		}
	case "high-card":
		if jokers == 1 {
			return ranks["one-pair"]
		} else {
			return ranks[label]
		}
	default:
		return ranks[label]
	}
}


type Hand struct {
	hand string
	bid  int
}

func (h *Hand) rank(part int) int {
	cards := strings.Split(h.hand, "")
	res := map[string]int{}
	for _, card := range cards {
		if count, ok := res[card]; ok {
			res[card] = count + 1
		} else {
			res[card] = 1
		}
	}
	counts := []int{}
	for _, count := range res {
		counts = append(counts, count)
	}
	sort.Ints(counts)
	if part == 1 {
		if reflect.DeepEqual(counts, []int{5}) {
			return ranks["five-of-a-kind"]
		} else if reflect.DeepEqual(counts, []int{1,4}) {
			return ranks["four-of-a-kind"]
		} else if reflect.DeepEqual(counts, []int{2,3}) {
			return ranks["full-house"]
		} else if reflect.DeepEqual(counts, []int{1,1,3}) {
			return ranks["three-of-a-kind"]
		} else if reflect.DeepEqual(counts, []int{1,2,2}) {
			return ranks["two-pair"]
		} else if reflect.DeepEqual(counts, []int{1,1,1,2}) {
			return ranks["one-pair"]
		} else {
			return ranks["high-card"]
		}
	} else {
		jCount := 0
		if jokers, ok := res["J"]; ok {
			jCount = jokers
		}
		if reflect.DeepEqual(counts, []int{5}) {
			return upgrade("five-of-a-kind", jCount)
		} else if reflect.DeepEqual(counts, []int{1,4}) {
			return upgrade("four-of-a-kind", jCount)
		} else if reflect.DeepEqual(counts, []int{2,3}) {
			return upgrade("full-house", jCount)
		} else if reflect.DeepEqual(counts, []int{1,1,3}) {
			return upgrade("three-of-a-kind", jCount)
		} else if reflect.DeepEqual(counts, []int{1,2,2}) {
			return upgrade("two-pair", jCount)
		} else if reflect.DeepEqual(counts, []int{1,1,1,2}) {
			return upgrade("one-pair", jCount)
		} else {
			return upgrade("high-card", jCount)
		}
	}
}

func (h *Hand) higher(other *Hand, part int) bool {
	log.Infof("Comparing %v to %v", h, other)
	if h.rank(part) > other.rank(part) {
		log.Info("left hand bigger")
		return true
	} else if h.rank(part) < other.rank(part) {
		log.Info("right hand bigger")
		return false
	} else {
		// equal check individual cards
		leftCards := strings.Split(h.hand, "")
		rightCards := strings.Split(other.hand, "")
		if part == 1 {
			for i, card := range leftCards {
				if cardRanks[card] > cardRanks[rightCards[i]] {
					return true
				} else if cardRanks[card] < cardRanks[rightCards[i]] {
					return false
				}
			}
		} else {
			for i, card := range leftCards {
				if jokerRanks[card] > jokerRanks[rightCards[i]] {
					return true
				} else if jokerRanks[card] < jokerRanks[rightCards[i]] {
					return false
				}
			}
		}
		// do we ever get exactly equal hands?
		return true
	}
}

func Day7(input *bufio.Scanner, part int) (result int) {
	hands := []*Hand{}
	for input.Scan() {
		text := strings.Split(input.Text(), " ")
		hand := text[0]
		bid, _ := strconv.Atoi(text[1])
		hands = append(hands, &Hand{ hand: hand, bid: bid })
	}

	sorted := sortByRank(hands, part)

	for _, h := range sorted {
		log.Info(h)
	}
	result = scoreHands(sorted)

	return
}

func sortByRank(hands []*Hand, part int) []*Hand {
	sorted := []*Hand{}
	for _, hand := range hands {
		if len(sorted) == 0 {
			log.Info("inserting first hand")
			sorted = append(sorted, hand)
		} else {
			for i, sortHand := range sorted {
				if hand.higher(sortHand, part) {
					sorted = append(sorted[:i + 1], sorted[i:]...)
					sorted[i] = hand
					break
				} else {
					if i == len(sorted) - 1 {
						sorted = append(sorted, hand)
					}
				}
			}
		}
	}
	return sorted
}

func scoreHands(hands []*Hand) int {
	highScore := len(hands)
	total := 0
	for _, hand := range hands {
		total += hand.bid * highScore
		highScore--
	}
	return total
}
