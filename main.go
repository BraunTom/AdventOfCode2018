package main

import (
	"AdventOfCode2018/FifthDay"
	"AdventOfCode2018/FirstDay"
	"AdventOfCode2018/FourthDay"
	"AdventOfCode2018/SecondDay"
	"AdventOfCode2018/SeventhDay"
	"AdventOfCode2018/ThirdDay"
	"fmt"
)

func main() {
	day := 7

	switch day {
	case 1:
		fmt.Println(FirstDay.Frequency())
		fmt.Println(FirstDay.Twice())
	case 2:
		fmt.Println(SecondDay.Checksum())
		fmt.Println(SecondDay.OneApart())
	case 3:
		fmt.Println(ThirdDay.Overlap())
		fmt.Println(ThirdDay.NoOverlap())
	case 4:
		fmt.Println(FourthDay.SleepTime())
	case 5:
		fmt.Println(FifthDay.Reduction())
		fmt.Println(FifthDay.ShortestPolymer())
	case 6:

	case 7:
		fmt.Println(SeventhDay.Order())
		fmt.Println(SeventhDay.Duration())
	}
}
