package day2

import (
	"AdventOfCode2022/src"
	"fmt"
)

const Win = 6
const Draw = 3
const Loss = 0

const Rock = 1
const Paper = 2
const Scissors = 3

func Run() {
	part1()
	part2()
}

func part1() {
	lines := src.ReadPuzzelInput("PuzzelInput/0201.txt")

	score := 0
	for _, m := range lines {
		if m == "" {
			break
		}
		tmp := 0
		switch m[0] {
		case 'A':
			switch m[2] {
			case 'X':
				tmp += Draw + Rock
				break
			case 'Y':
				tmp += Win + Paper
				break
			case 'Z':
				tmp += Loss + Scissors
				break
			}
			break
		case 'B':
			switch m[2] {
			case 'X':
				tmp += Loss + Rock
				break
			case 'Y':
				tmp += Draw + Paper
				break
			case 'Z':
				tmp += Win + Scissors
				break
			}
			break
		case 'C':
			switch m[2] {
			case 'X':
				tmp += Win + Rock
				break
			case 'Y':
				tmp += Loss + Paper
				break
			case 'Z':
				tmp += Draw + Scissors
				break
			}
			break
		}

		score += tmp
	}

	fmt.Println("Part 1 solution", score)
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0201.txt")

	score := 0
	for _, m := range lines {
		if m == "" {
			break
		}
		tmp := 0
		switch m[0] {
		case 'A':
			switch m[2] {
			case 'X':
				tmp += Loss + Scissors
				break
			case 'Y':
				tmp += Draw + Rock
				break
			case 'Z':
				tmp += Win + Paper
				break
			}
			break
		case 'B':
			switch m[2] {
			case 'X':
				tmp += Loss + Rock
				break
			case 'Y':
				tmp += Draw + Paper
				break
			case 'Z':
				tmp += Win + Scissors
				break
			}
			break
		case 'C':
			switch m[2] {
			case 'X':
				tmp += Loss + Paper
				break
			case 'Y':
				tmp += Draw + Scissors
				break
			case 'Z':
				tmp += Win + Rock
				break
			}
			break
		}

		score += tmp
	}

	fmt.Println("Part 2 solution", score)
}
