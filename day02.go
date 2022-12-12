package main

var oponentTrick = map[byte]int{'A': 1, 'B': 2, 'C': 3}
var myTrick = map[byte]int{'X': 1, 'Y': 2, 'Z': 3}
var losingOrder = map[int]int{1: 3, 2: 1, 3: 2}
var winingOrder = map[int]int{3: 1, 1: 2, 2: 3}

func Day02() {
	data := ReadInput("input/day02.txt")

	println("Day 02 - Part 01: ", part01(data))
	println("Day 02 - Part 01: ", part02(data))
}

func part01(data []string) int {
	var score int

	for _, round := range data {
		if oponentTrick[round[0]] == myTrick[round[2]] { // Draw
			score += 3
		} else {
			if losingOrder[oponentTrick[round[0]]] == myTrick[round[2]] { // Loss
				score += 0
			} else { // Win
				score += 6
			}
		}
		score += myTrick[round[2]]
	}

	return score
}

func part02(data []string) int {
	var score int

	for _, round := range data {
		if round[2] == 'X' { // Loss
			score += 0
			score += losingOrder[oponentTrick[round[0]]]
		} else if round[2] == 'Y' { // Draw
			score += 3
			score += oponentTrick[round[0]]
		} else if round[2] == 'Z' { // Win
			score += 6
			score += winingOrder[oponentTrick[round[0]]]
		}
	}

	return score
}
