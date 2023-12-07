package calendar

import (
	"bufio"
	"strconv"
	"strings"
)

func init() {
	registry["Day5"] = Day5
}

// var seeds = []int{ 79, 14, 55, 13 }
var seeds = []int{ 919339981, 562444630, 3366006921, 67827214, 1496677366, 101156779, 4140591657, 5858311, 2566406753, 71724353, 2721360939, 35899538, 383860877, 424668759, 3649554897, 442182562, 2846055542, 49953829, 2988140126, 256306471 }

func Day5(input *bufio.Scanner, part int) (result int) {

//	input = bufio.NewScanner(strings.NewReader(`seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4
// `))

	collectRange := false
	var rangeKey string
	rangeMap := map[string][][]string{}

	result = 1000000000000000

	seedRanges := [][]int{}
	for i, seed := range seeds {
		if i%2 == 0 {
			seedRanges = append(seedRanges, []int{seed})
		} else {
			seedRanges[i/2] = append(seedRanges[i/2], seedRanges[i/2][0] + seed - 1)
		}
	}
	log.Info(seedRanges)
	for input.Scan() {
		if !collectRange {
			if strings.Contains(input.Text(), "seed-to-soil") {
				collectRange = true
				rangeKey = "seed-to-soil"
			} else if strings.Contains(input.Text(), "soil-to-fertilizer") {
				collectRange = true
				rangeKey = "soil-to-fertilizer"
			} else if strings.Contains(input.Text(), "fertilizer-to-water") {
				collectRange = true
				rangeKey = "fertilizer-to-water"
			} else if strings.Contains(input.Text(), "water-to-light") {
				collectRange = true
				rangeKey = "water-to-light"
			} else if strings.Contains(input.Text(), "light-to-temperature") {
				collectRange = true
				rangeKey = "light-to-temperature"
			} else if strings.Contains(input.Text(), "temperature-to-humidity") {
				collectRange = true
				rangeKey = "temperature-to-humidity"
			} else if strings.Contains(input.Text(), "humidity-to-location") {
				collectRange = true
				rangeKey = "humidity-to-location"
			}
		} else {
			ranges := strings.Split(input.Text(), " ")
			if len(ranges) == 1 {
				collectRange = false
			} else {
				rangeMap[rangeKey] = append(rangeMap[rangeKey], ranges)
			}
		}
	}

	log.Debug(rangeMap)
	if part == 1 {
		for _, seed := range seeds {
			log.Infof("Mapping seed %d to soil", seed)
			for _, rang := range rangeMap["seed-to-soil"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping soil %d to fertilizer", seed)
			for _, rang := range rangeMap["soil-to-fertilizer"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping fertilizer %d to water", seed)
			for _, rang := range rangeMap["fertilizer-to-water"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping water %d to light", seed)
			for _, rang := range rangeMap["water-to-light"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping light %d to temperature", seed)
			for _, rang := range rangeMap["light-to-temperature"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping temperature %d to humidity", seed)
			for _, rang := range rangeMap["temperature-to-humidity"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Mapping humidity %d to location", seed)
			for _, rang := range rangeMap["humidity-to-location"] {
				dstRangeStart, _ := strconv.Atoi(rang[0])
				srcRangeStart, _ := strconv.Atoi(rang[1])
				length, _ := strconv.Atoi(rang[2])
				// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
				if seed >= srcRangeStart && seed < (srcRangeStart + length) {
					seed = dstRangeStart - srcRangeStart + seed
					break
				}
			}

			// log.Infof("Final location %d", seed)
			if seed < result {
				result = seed
			}
		}

		return
	} else {
		for _, seedRange := range seedRanges {
			for s := seedRange[0]; s <= seedRange[1]; s++ {
				seed := s
				// log.Infof("Mapping seed %d to soil", seed)
				for _, rang := range rangeMap["seed-to-soil"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping soil %d to fertilizer", seed)
				for _, rang := range rangeMap["soil-to-fertilizer"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping fertilizer %d to water", seed)
				for _, rang := range rangeMap["fertilizer-to-water"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping water %d to light", seed)
				for _, rang := range rangeMap["water-to-light"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping light %d to temperature", seed)
				for _, rang := range rangeMap["light-to-temperature"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping temperature %d to humidity", seed)
				for _, rang := range rangeMap["temperature-to-humidity"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Mapping humidity %d to location", seed)
				for _, rang := range rangeMap["humidity-to-location"] {
					dstRangeStart, _ := strconv.Atoi(rang[0])
					srcRangeStart, _ := strconv.Atoi(rang[1])
					length, _ := strconv.Atoi(rang[2])
					// log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length)
					if seed >= srcRangeStart && seed < (srcRangeStart + length) {
						seed = dstRangeStart - srcRangeStart + seed
						break
					}
				}

				// log.Infof("Final location %d", seed)
				if seed < result {
					result = seed
				}
			}
		}
		return result
		// for _, sRange := range seedRanges {
		//	log.Infof("Mapping seed range %v to soil", sRange)
		//	next := compareRanges(sRange, rangeMap["seed-to-soil"])
		//	log.Infof("Soil ranges to process %v", next)

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["soil-to-fertilizer"])
		//		log.Infof("Fertilizer ranges to process %v", next)
		//	}

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["fertilizer-to-water"])
		//		log.Infof("Water ranges to process %v", next)
		//	}

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["water-to-light"])
		//		log.Infof("Light ranges to process %v", next)
		//	}

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["light-to-temperature"])
		//		log.Infof("Temperature ranges to process %v", next)
		//	}

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["temperature-to-humidity"])
		//		log.Infof("Humidity ranges to process %v", next)
		//	}

		//	for _, sr := range next {
		//		next = compareRanges(sr, rangeMap["humidity-to-location"])
		//		log.Infof("Location ranges %v", next)
		//		for _, r := range next {
		//			if r[0] < result {
		//				result = r[0]
		//			}
		//		}
		//	}
		// }
	}
}

func compareRanges(src []int, dst [][]string) [][]int {
	results := [][]int{}
	leftovers := [][]int{src}
	for _, rang := range dst {
		log.Debug("looping over ranges")
		if len(leftovers) > 0 {
			log.Debugf("Leftovers to process %v", leftovers)
			pop := leftovers[0]
			leftovers = leftovers[1:]
			start := pop[0]
			end := pop[1]

			dstRangeStart, _ := strconv.Atoi(rang[0])
			srcRangeStart, _ := strconv.Atoi(rang[1])
			length, _ := strconv.Atoi(rang[2])
			transform := dstRangeStart - srcRangeStart

			log.Infof("checking if between %d and %d", srcRangeStart, srcRangeStart + length - 1)
			if end < srcRangeStart || start >= (srcRangeStart + length) {
				log.Debug("no overlap, pass through")
				leftovers = append(leftovers, []int{ start, end })
			} else if start >= srcRangeStart && end < (srcRangeStart + length) {
				log.Debug("full overlap, transform")
				results = append(results, []int{ start + transform, end + transform })
				log.Debugf("current results %v", results)
				log.Debugf("current leftovers %v", leftovers)
				break
			} else {
				if srcRangeStart >= start {
					log.Debug("slicing off front")
					newEnd := srcRangeStart - start
					leftovers = append(leftovers, []int{start, newEnd})
					start = srcRangeStart
				}
				if end >= srcRangeStart + length - 1 {
					log.Debug("slicing off back")
					newStart := srcRangeStart + length
					leftovers = append(leftovers, []int{newStart, end})
					end = srcRangeStart + length - 1
				}
				log.Debug("transform what is left")
				results = append(results, []int{ start + transform, end + transform })
			}
			log.Infof("leftovers are %v", leftovers)
		}
	}
	if len(leftovers) > 0 {
		log.Debugf("appending leftovers %v", leftovers)
		results = append(results, leftovers...)
	}
	log.Debugf("returning results %v", results)
	return results
}
