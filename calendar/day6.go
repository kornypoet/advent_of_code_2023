package calendar

import (
        "bufio"
	"math"
)

func init() {
        registry["Day6"] = Day6
}

func Day6(input *bufio.Scanner, part int) (result int) {
	var answer float64 = 1
	var parameters map[float64]float64
	if part == 1 {
		parameters = map[float64]float64{
			56: 499,
			97: 2210,
			77: 1097,
			93: 1440,
		}
	} else {
		parameters = map[float64]float64{
			56977793: 499221010971440,
		}
	}
	for time, dist := range parameters {
		answer *= diffRoots(time, dist)
	}
	
        return int(answer)
}

func diffRoots(time float64, dist float64) float64 {	
	discriminant := math.Sqrt((time * time) - (4 * dist)) / -2
	rootOne := (time / 2) + discriminant
	rootTwo := (time / 2) - discriminant
	return math.Floor(rootTwo) - math.Ceil(rootOne) + 1
}
