package main

import (
	"AdventOfCode2022/src/day1"
	"AdventOfCode2022/src/day2"
	"AdventOfCode2022/src/day3"
	"AdventOfCode2022/src/day4"
	"AdventOfCode2022/src/day5"
	"AdventOfCode2022/src/day6"
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
	case 2:
		day2.Run()
		break
	case 3:
		day3.Run()
		break
	case 4:
		day4.Run()
		break
	case 5:
		day5.Run()
		break
	case 6:
		day6.Run()
		break
	}
}
