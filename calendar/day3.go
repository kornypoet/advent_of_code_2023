package calendar

import (
	"bufio"
	"math"
	"strconv"
	"unicode"
)

func init() {
	registry["Day3"] = Day3
}

type Symbol struct {
	char        string
	adjacencies []*EnginePart
	location    []int
}

func (s *Symbol) isGear() bool {
	return len(s.adjacencies) == 2 && s.char == "*"
}

func (s *Symbol) ratio() int {
	left, _ := strconv.Atoi(s.adjacencies[0].number)
	right, _ := strconv.Atoi(s.adjacencies[1].number)
	return left * right
}

func (s *Symbol) checkAdj(e *EnginePart) {
	checked := false
	for _, loc := range e.locations {
		if !checked {
			if math.Abs(float64(loc[0] - s.location[0])) <= 1 && math.Abs(float64(loc[1] - s.location[1])) <= 1 {
				s.adjacencies = append(s.adjacencies, e)
				checked = true
			}
		}
	}
}

type EnginePart struct {
	collect   bool
	number    string
	locations [][]int
	adjacent  bool
}

func NewEnginePart() *EnginePart {
	return &EnginePart{}
}

func (e *EnginePart) checkAdj(sym *Symbol) {
	for _, loc := range e.locations {
		if math.Abs(float64(loc[0] - sym.location[0])) <= 1 && math.Abs(float64(loc[1] - sym.location[1])) <= 1 {
			e.adjacent = true
		}
	}
}

func Day3(input *bufio.Scanner, part int) (result int) {
	rowPos := 0
	parts := []*EnginePart{}
	symbols := []*Symbol{}

	for input.Scan() {
		chars := []rune(input.Text())
		engPart := NewEnginePart()
		for colPos, char := range chars {
			if unicode.IsDigit(char) {
				if !engPart.collect {
					engPart.collect = true
					parts = append(parts, engPart)
				}
				engPart.number += string(char)
				engPart.locations = append(engPart.locations, []int{ colPos, rowPos })
			} else {
				if engPart.collect {
					engPart.collect = false
					log.Debug(engPart)
					engPart = NewEnginePart()
				}
				if char != '.' {
					sym := &Symbol{ char: string(char), location: []int{ colPos, rowPos }}
					log.Debug(sym)
					symbols = append(symbols, sym)
				}
			}
		}
		rowPos++
	}

	for _, engPart := range parts {
		// could do this better. Only need to check rows above, below and current
		for _, sym := range symbols {
			engPart.checkAdj(sym)
		}
		if engPart.adjacent {
			num, _ := strconv.Atoi(engPart.number)
			result += num
		} else {
			log.Debug(engPart)
		}
	}
	if part == 1 {
		return
	} else {
		result = 0
		for _, sym := range symbols {
			for _, engPart := range parts {
				sym.checkAdj(engPart)
			}
			log.Debug(sym)
			if sym.isGear() {
				result += sym.ratio()
			}
		}
		return
	}
}
