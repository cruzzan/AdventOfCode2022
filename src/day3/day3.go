package day3

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
	lines := src.ReadPuzzelInput("PuzzelInput/0301.txt")

	score := 0
	for _, l := range lines {
		com1 := l[:len(l)/2]
		com2 := l[len(l)/2:]
		for _, c := range uniqueRunes(com1) {
			packSum := 0
			covered := make([]rune, 0)
			if strings.ContainsRune(com2, c) && !contains(covered, c) {
				covered = append(covered, c)
				packSum += itemPriority(c)
			}
			score += packSum
		}
	}

	fmt.Println("Part 1 solution", score)
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0301.txt")

	score := 0
	for i := 0; i < len(lines); i += 3 {
		p1 := uniqueRunes(lines[i])
		p2 := uniqueRunes(lines[i+1])
		p3 := uniqueRunes(lines[i+2])

		combined := p1 + p2 + p3

		for _, c := range combined {
			if strings.Count(combined, string(c)) >= 3 {
				score += itemPriority(c)
				break
			}
		}
	}

	fmt.Println("Part 2 solution", score)
}

func itemPriority(c rune) int {
	if c >= 65 && c <= 90 {
		return int(c) - 38
	} else if c >= 97 && c <= 122 {
		return int(c) - 96
	}

	return 0
}

func contains(haystack []rune, needle rune) bool {
	for _, c := range haystack {
		if c == needle {
			return true
		}
	}
	return false
}

func uniqueRunes(pile string) string {
	unique := make([]rune, 0)

	for _, c := range pile {
		if !strings.ContainsRune(string(unique), c) {
			unique = append(unique, c)
		}
	}

	return string(unique)
}
