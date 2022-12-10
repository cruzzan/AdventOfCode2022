package day8

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
	lines := src.ReadPuzzelInput("PuzzelInput/0801.txt")
	maptrix := buildMapMatrix(lines)

	visible := 0
	for i := 1; i < len(maptrix)-1; i++ {
		for j := 1; j < len(maptrix[0])-1; j++ {
			if isVisible(maptrix, i, j) {
				visible++
			}
		}
	}

	fmt.Println("Part 1 solution", visible+(98*4))
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0801.txt")
	maptrix := buildMapMatrix(lines)

	mostScenic := 0
	for i := 1; i < len(maptrix)-1; i++ {
		for j := 1; j < len(maptrix[0])-1; j++ {
			ss := getScenicScore(maptrix, i, j)
			if ss > mostScenic {
				mostScenic = ss
			}
		}
	}

	fmt.Println("Part 2 solution", mostScenic)
}

func getScenicScore(maptrix [][]int, i, j int) int {
	_, n := north(maptrix, i, j)
	_, s := south(maptrix, i, j)
	_, e := east(maptrix, i, j)
	_, w := west(maptrix, i, j)

	return n * s * e * w
}

func isVisible(maptrix [][]int, i, j int) bool {
	n, _ := north(maptrix, i, j)
	s, _ := south(maptrix, i, j)
	e, _ := east(maptrix, i, j)
	w, _ := west(maptrix, i, j)

	return n || s || e || w
}

func north(maptrix [][]int, i, j int) (bool, int) {
	viewDistance := 0
	for k := i - 1; k >= 0; k-- {
		viewDistance++
		if maptrix[i][j] <= maptrix[k][j] {
			return false, viewDistance
		}
	}
	return true, viewDistance
}

func south(maptrix [][]int, i, j int) (bool, int) {
	viewDistance := 0
	for k := i + 1; k < len(maptrix[0]); k++ {
		viewDistance++
		if maptrix[i][j] <= maptrix[k][j] {
			return false, viewDistance
		}
	}
	return true, viewDistance
}

func east(maptrix [][]int, i, j int) (bool, int) {
	viewDistance := 0
	for k := j + 1; k < len(maptrix[0]); k++ {
		viewDistance++
		if maptrix[i][j] <= maptrix[i][k] {
			return false, viewDistance
		}
	}
	return true, viewDistance
}

func west(maptrix [][]int, i, j int) (bool, int) {
	viewDistance := 0
	for k := j - 1; k >= 0; k-- {
		viewDistance++
		if maptrix[i][j] <= maptrix[i][k] {
			return false, viewDistance
		}
	}
	return true, viewDistance
}

func buildMapMatrix(lines []string) [][]int {
	matrix := make([][]int, 0)

	for _, l := range lines {
		items := strings.Split(l, "")

		row := make([]int, 0)
		for _, item := range items {
			v, err := strconv.Atoi(item)
			if err != nil {
				fmt.Println(err)
			}

			row = append(row, v)
		}

		matrix = append(matrix, row)
	}

	return matrix
}
