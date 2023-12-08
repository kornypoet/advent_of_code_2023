package calendar

import (
	"bufio"
	"regexp"
	"strings"
)

func init() {
	registry["Day8"] = Day8
}

var directions = `LRLRRLRLRLLRRRLLLRLLRRLLRRRLRLRLRRRLRRLRLRRRLRRRLRRRLRRRLRRLRRRLRRRLRRLLLRLLRLRRRLRRRLRRLRLRLRLRRRLRRRLRRRLRRLRRLRRLLRRLRRRLLRRLRRRLRRRLRRRLRLRRLRLRRRLRRLLRLRLLRLRLRRRLRLRRLLRRRLRLRLRLRLRLRRLRLRRLLLLRRLRRLRRRLRRLRRLRRRLRRLRRRLLRLRRLLRLRRLRRLRRLLRRRLRLRLRRRLRRLRLLRLRRRR`
// var directions = `RL`
// var directions = `LLR`
// var directions = `LR`

type Node struct {
	label string
	left  *Node
	right *Node
}

var nodeGraph = map[string]*Node{}

func Day8(input *bufio.Scanner, part int) (result int) {
	re := regexp.MustCompile(`(?P<node>\w{3}) = \((?P<left>\w{3}), (?P<right>\w{3})\)`)
	for input.Scan() {
		match := re.FindStringSubmatch(input.Text())
		label := match[1]
		left := match[2]
		right := match[3]

		var node *Node
		node, ok := nodeGraph[label]
		if !ok {
			node = &Node{ label: label }
			nodeGraph[label] = node
		}

		if leftNode, ok := nodeGraph[left]; ok {
			node.left = leftNode
		} else {
			ln := &Node{ label: left }
			node.left = ln
			nodeGraph[left] = ln
		}

		if rightNode, ok := nodeGraph[right]; ok {
			node.right = rightNode
		} else {
			rn := &Node{ label: right }
			node.right = rn
			nodeGraph[right] = rn
		}
	}

	if part == 1 {
		steps := strings.Split(directions, "")
		step := 0
		zNodeFound := false
		currNode := nodeGraph["AAA"]

		for {
			// log.Infof("Current Node %#v", currNode)
			// log.Infof("Current step %d", step)
			direction := steps[step % len(steps)]
			step++
			switch direction {
			case "L":
				currNode = currNode.left
			case "R":
				currNode = currNode.right
			}
			zNodeFound = currNode.label == "ZZZ"

			if zNodeFound {
				break
			}
		}

		return step
	} else {
		currNodes := []*Node{}
		for k, node := range nodeGraph {
			if strings.HasSuffix(k, "A") {
				currNodes = append(currNodes, node)
			}
		}

		steps := strings.Split(directions, "")
		results := []int{}

		for _, node := range currNodes {
			currNode := node
			zNodeFound := false
			step := 0
			log.Infof("Current node %#v", currNode)

			for {
				direction := steps[step % len(steps)]
				step++
				switch direction {
				case "L":
					currNode = currNode.left
				case "R":
					currNode = currNode.right
				}

				if strings.HasSuffix(currNode.label, "Z") {
					zNodeFound = true
				}

				if zNodeFound {
					log.Infof("Step Count %d", step)
					results = append(results, step)
					break
				}
			}
		}		
		return LCM(results[0], results[1], results[2:]...)
	}
}

func GCD(a, b int) int {
	for b != 0 {
		t := b
		b = a % b
		a = t
	}
	return a
}

func LCM(a, b int, integers ...int) int {
	result := a * b / GCD(a, b)

	for i := 0; i < len(integers); i++ {
		result = LCM(result, integers[i])
	}

	return result
}
