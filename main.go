package main

import (
	"AdventOfCode2022/src/day1"
	"fmt"
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	day = kingpin.Flag("day", "The day you want to get the result for.").Short('d').Int()
)

func main() {
	kingpin.Parse()
	fmt.Println("Hello there, and welcome back for 2022")
	fmt.Println("Running day", *day)

	switch *day {
	case 1:
		day1.Run()
		break
	}
}
