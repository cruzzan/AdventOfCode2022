package day1

import (
	"AdventOfCode2022/src"
	"fmt"
	"os"
	"strconv"
)

func Run() {
	part1()
	part2()
}

func part1() {
	elf := make(map[int]int, 0)
	in := src.ReadPuzzelInput("PuzzelInput/0101.txt")

	elfCounter := 0
	for _, l := range in {
		if l == "" {
			elfCounter++
			continue
		} else {
			cal, err := strconv.Atoi(l)
			if err != nil {
				fmt.Println("Something went wrong", err)
				os.Exit(1)
			}

			elf[elfCounter] += cal
		}
	}

	highest := 0
	for _, cal := range elf {
		if cal > highest {
			highest = cal
		}
	}

	fmt.Println("Part 1 solution", highest)
}

func part2() {
	elf := make(map[int]int, 0)
	in := src.ReadPuzzelInput("PuzzelInput/0101.txt")

	elfCounter := 0
	for _, l := range in {
		if l == "" {
			elfCounter++
			continue
		} else {
			cal, err := strconv.Atoi(l)
			if err != nil {
				fmt.Println("Something went wrong", err)
				os.Exit(1)
			}

			elf[elfCounter] += cal
		}
	}

	first := 0
	second := 0
	third := 0
	for _, cal := range elf {
		if cal > first {
			third = second
			second = first
			first = cal
		} else if cal > second {
			third = second
			second = cal
		} else if cal > third {
			third = cal
		}
	}

	fmt.Println("Part 2 solution", first+second+third)
}
