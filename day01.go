package main

import (
	"sort"
	"strconv"
)

func Day01() {
	var data = ReadInput("input/day01.txt")
	var calories []int
	var buffer, count int

	for _, line := range data {
		if line == "" {
			calories = append(calories, count)
			count = 0
		} else {
			buffer, _ = strconv.Atoi(line)
			count += buffer
		}
	}

	sort.Ints(calories)

	println("Day 01 - Part 01: ", calories[len(calories)-1])
	println("Day 01 - Part 02: ", calories[len(calories)-3]+calories[len(calories)-2]+calories[len(calories)-1])
}
