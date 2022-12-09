package day7

import (
	"AdventOfCode2022/src"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

const FILE = 0
const DIR = 1

type node struct {
	name     string
	fType    int
	size     int
	children []*node
}

func (n *node) insert(path []string, new *node) {
	if len(path) > 1 {
		for _, c := range n.children {
			if c.name == path[0] {
				c.insert(path[1:], new)
				return
			}
		}
		d := node{
			name:  path[0],
			fType: DIR,
			size:  0,
		}

		d.insert(path[1:], new)
		n.children = append(n.children, &d)
	} else {
		n.children = append(n.children, new)
	}
}

func (n *node) propagateSize() int {
	for _, c := range n.children {
		n.size += c.propagateSize()
	}

	return n.size
}

func (n *node) sumOver100K() int {
	sum := 0

	for _, c := range n.children {
		sum += c.sumOver100K()
	}

	if n.fType == DIR && n.size < 100000 {
		return sum + n.size
	}
	return sum
}

func (n *node) smallestAbove(min, smallest int) int {
	for _, c := range n.children {
		candidate := c.smallestAbove(min, smallest)
		if candidate > min && candidate < smallest {
			smallest = candidate
		}
	}

	if n.size > min && n.size < smallest {
		return n.size
	}
	return smallest
}

func Run() {
	part1()
	part2()
}

func part1() {
	lines := src.ReadPuzzelInput("PuzzelInput/0701.txt")

	cursor := ""
	root := node{
		name:     "/",
		size:     0,
		children: nil,
	}

	for _, l := range lines {
		if l[:1] == "$" {
			pattern := regexp.MustCompile(`\$\s(\w+)\s*(.*)`)
			m := pattern.FindStringSubmatch(l)
			if m[1] == "cd" {
				if m[2] == "/" {
					cursor = ""
				} else if m[2] == ".." {
					i := strings.LastIndex(cursor, "/")
					cursor = cursor[:i]
				} else {
					cursor = fmt.Sprintf("%s/%s", cursor, m[2])
				}
			}
		} else {
			pattern := regexp.MustCompile(`(\d+)\s(.+)`)
			m := pattern.FindStringSubmatch(l)

			if len(m) > 0 {
				size, err := strconv.Atoi(m[1])
				if err != nil {
					fmt.Println("Could not convert string to int", err)
				}

				fileName := fmt.Sprintf("%s/%s", cursor, m[2])
				f := strings.Split(fileName, "/")

				n := node{
					name:  fileName,
					fType: FILE,
					size:  size,
				}
				root.insert(f[1:], &n)
			}
		}
	}

	root.propagateSize()

	fmt.Println("Part 1 solution", root.size, root.sumOver100K())
}

func part2() {
	lines := src.ReadPuzzelInput("PuzzelInput/0701.txt")

	cursor := ""
	root := node{
		name:     "/",
		size:     0,
		children: nil,
	}

	for _, l := range lines {
		if l[:1] == "$" {
			pattern := regexp.MustCompile(`\$\s(\w+)\s*(.*)`)
			m := pattern.FindStringSubmatch(l)
			if m[1] == "cd" {
				if m[2] == "/" {
					cursor = ""
				} else if m[2] == ".." {
					i := strings.LastIndex(cursor, "/")
					cursor = cursor[:i]
				} else {
					cursor = fmt.Sprintf("%s/%s", cursor, m[2])
				}
			}
		} else {
			pattern := regexp.MustCompile(`(\d+)\s(.+)`)
			m := pattern.FindStringSubmatch(l)

			if len(m) > 0 {
				size, err := strconv.Atoi(m[1])
				if err != nil {
					fmt.Println("Could not convert string to int", err)
				}

				fileName := fmt.Sprintf("%s/%s", cursor, m[2])
				f := strings.Split(fileName, "/")

				n := node{
					name:  fileName,
					fType: FILE,
					size:  size,
				}
				root.insert(f[1:], &n)
			}
		}
	}

	root.propagateSize()

	unused := 70000000 - root.size
	min := 30000000 - unused

	fmt.Println("Part 2 solution", root.smallestAbove(min, root.size))
}
