package day5

import (
	"AdventOfCode2022/src"
	"fmt"
	"regexp"
	"strconv"
)

func Run() {
	part1()
	part2()
}

func part1() {
	lines := src.ReadPuzzelInput("PuzzelInput/0501.txt")
	lines = lines[10:]

	stacks := getStartingStacks()

	for _, l := range lines {
		count, from, to := extractInstructions(l)

		stacks = move(stacks, count, from, to, true)
	}

	tops := ""
	for i := 1; i <= 9; i++ {
		if len(stacks[i]) > 0 {
			tops = fmt.Sprintf("%s %s", tops, stacks[i][0])
		} else {
			tops = fmt.Sprintf("%s %s", tops, " ")
		}
	}

	fmt.Println("Part 1 solution", tops)
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0501.txt")
	lines = lines[10:]

	stacks := getStartingStacks()

	for _, l := range lines {
		count, from, to := extractInstructions(l)

		stacks = move(stacks, count, from, to, false)
	}

	tops := ""
	for i := 1; i <= 9; i++ {
		if len(stacks[i]) > 0 {
			tops = fmt.Sprintf("%s %s", tops, stacks[i][0])
		} else {
			tops = fmt.Sprintf("%s %s", tops, " ")
		}
	}

	fmt.Println("Part 2 solution", tops)
}

func extractInstructions(line string) (int, int, int) {
	pattern := regexp.MustCompile(`(\d+).*(\d+).*(\d+)`)
	m := pattern.FindStringSubmatch(line)

	count, err := strconv.Atoi(m[1])
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}
	from, err := strconv.Atoi(m[2])
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}
	to, err := strconv.Atoi(m[3])
	if err != nil {
		fmt.Println("Error converting string to int", err)
	}

	return count, from, to
}

func getStartingStacks() map[int][]string {
	stacks := make(map[int][]string)
	stacks[1] = []string{"V", "N", "F", "S", "M", "P", "H", "J"}
	stacks[2] = []string{"Q", "D", "J", "M", "L", "R", "S"}
	stacks[3] = []string{"B", "W", "S", "C", "H", "D", "Q", "N"}
	stacks[4] = []string{"L", "C", "S", "R"}
	stacks[5] = []string{"B", "F", "P", "T", "V", "M"}
	stacks[6] = []string{"C", "N", "Q", "R", "T"}
	stacks[7] = []string{"R", "V", "G"}
	stacks[8] = []string{"R", "L", "D", "P", "S", "Z", "C"}
	stacks[9] = []string{"F", "B", "P", "G", "V", "J", "S", "D"}

	return stacks
}

func move(stacks map[int][]string, count, from, to int, reverseMoved bool) map[int][]string {
	moving := make([]string, count)
	rest := make([]string, len(stacks[to]))
	copy(moving, stacks[from][:count])
	copy(rest, stacks[to][0:])

	if reverseMoved {
		reverse(moving)
	}

	stacks[to] = append(moving, rest...)
	stacks[from] = stacks[from][count:]

	return stacks
}

func reverse(input []string) {
	for i, j := 0, len(input)-1; i < j; i, j = i+1, j-1 {
		input[i], input[j] = input[j], input[i]
	}
}
