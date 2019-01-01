package main

import (
	"AdventOfCode2018/FirstDay"
	"AdventOfCode2018/SecondDay"
	"AdventOfCode2018/ThirdDay"
	"fmt"
)

func main() {
	day := 2

	switch day {
	case 1:
		fmt.Println(FirstDay.Frequency())
		fmt.Println(FirstDay.Twice())
	case 2:
		fmt.Println(SecondDay.Checksum())
		fmt.Println(SecondDay.OneApart())
	case 3:
		fmt.Println(ThirdDay.Overlap())
	}
}
