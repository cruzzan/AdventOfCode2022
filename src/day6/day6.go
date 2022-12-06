package day6

import (
	"AdventOfCode2022/src"
	"fmt"
	"strings"
)

func Run() {
	part1()
	part2()
}

func part1() {
	line := src.ReadPuzzelInput("PuzzelInput/0601.txt")[0]

	endOfSoPM := 0
	for i := 4; i < len(line); i++ {
		if check(line[i-4 : i]) {
			endOfSoPM = i
			break
		}
	}

	fmt.Println("Part 1 solution", endOfSoPM)
}

func part2() {
	line := src.ReadPuzzelInput("PuzzelInput/0601.txt")[0]

	endOfSoPM := 0
	for i := 14; i < len(line); i++ {
		if check(line[i-14 : i]) {
			endOfSoPM = i
			break
		}
	}

	fmt.Println("Part 2 solution", endOfSoPM)
}

func check(buffer string) bool {
	res := true
	for _, char := range buffer {
		if strings.Count(buffer, string(char)) > 1 {
			res = false
			break
		}
	}
	return res
}
