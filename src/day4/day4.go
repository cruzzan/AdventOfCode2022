package day4

import (
	"AdventOfCode2022/src"
	"fmt"
	"strconv"
	"strings"
)

func Run() {
	part1()
	part2()
}

func part1() {
	lines := src.ReadPuzzelInput("PuzzelInput/0401.txt")

	fullOverlaps := 0
	for _, l := range lines {
		ranges := strings.Split(l, ",")

		elf1 := getRange(strings.Split(ranges[0], "-"))
		elf2 := getRange(strings.Split(ranges[1], "-"))

		overlap := findOverlap(elf1, elf2)

		if len(overlap) == len(elf1) || len(overlap) == len(elf2) {
			fullOverlaps++
		}
	}

	fmt.Println("Part 1 solution", fullOverlaps)
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0401.txt")

	fullOverlaps := 0
	for _, l := range lines {
		ranges := strings.Split(l, ",")

		elf1 := getRange(strings.Split(ranges[0], "-"))
		elf2 := getRange(strings.Split(ranges[1], "-"))

		overlap := findOverlap(elf1, elf2)

		if len(overlap) > 0 {
			fullOverlaps++
		}
	}

	fmt.Println("Part 2 solution", fullOverlaps)
}

func getRange(rangeLimits []string) []int {
	r := make([]int, 0)
	start, err := strconv.Atoi(rangeLimits[0])
	if err != nil {
		fmt.Println("Got error in getRange", err)
	}
	end, err := strconv.Atoi(rangeLimits[1])
	if err != nil {
		fmt.Println("Got error in getRange", err)
	}

	for i := start; i <= end; i++ {
		r = append(r, i)
	}

	return r
}

func findOverlap(p1 []int, p2 []int) []int {
	overlapp := make([]int, 0)

	for _, i := range p1 {
		for _, j := range p2 {
			if i == j {
				overlapp = append(overlapp, i)
				continue
			}
		}
	}

	return overlapp
}
