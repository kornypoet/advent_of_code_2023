package calendar

import (
	"bufio"
	"regexp"
	"strconv"
	"strings"
)

func init() {
	registry["Day2"] = Day2
}

var maximum = map[string]int{
	"red": 12,
	"green": 13,
	"blue": 14,
}

func Day2(input *bufio.Scanner, part int) (result int) {
	games := make(map[int][]map[string]int)
	for input.Scan() {
		info := strings.Split(input.Text(), ":")
		re := regexp.MustCompile(`^Game (\d+)`)
		gameInfo := re.FindStringSubmatch(info[0])
		gameId, _ := strconv.Atoi(gameInfo[1])
		revealed := strings.Split(info[1], ";")
		games[gameId] = []map[string]int{}
		for _, set := range revealed {
			drawn := make(map[string]int)
			color := regexp.MustCompile(`(?P<count>\d+) (?P<color>blue|green|red)`)
			colorInfo := color.FindAllStringSubmatch(set, -1)
			for _, selected := range colorInfo {
				count, _ := strconv.Atoi(selected[1])
				drawn[selected[2]] = count
			}
			games[gameId] = append(games[gameId], drawn)
			log.Debug(drawn)
		}
	}
	if part == 1 {
		for id, drawings := range games {
			possible := true
			for _, drawn := range drawings {
				if drawn["red"] > maximum["red"] || drawn["green"] > maximum["green"] || drawn["blue"] > maximum["blue"] {
					possible = false
					log.Debug("Impossible")
					log.Debug(id)
					log.Debug(drawn)
					break
				}
			}
			if possible {
				result += id
			}
		}
		return
	} else {
		for _, drawings := range games {
			powerSet := map[string]int{
				"red": 0,
				"green": 0,
				"blue": 0,
			}
			for _, drawn := range drawings {
				for _, color := range []string{"red", "green", "blue"} {
					if drawn[color] > powerSet[color] {
						powerSet[color] = drawn[color]
					}
				}
			}
			power := powerSet["red"] * powerSet["green"] * powerSet["blue"]
			result += power
		}
		return
	}
}
