package main

import (
	"AdventOfCode2018/EightDay"
	"AdventOfCode2018/EleventhDay"
	"AdventOfCode2018/FifthDay"
	"AdventOfCode2018/FirstDay"
	"AdventOfCode2018/FourteenthDay"
	"AdventOfCode2018/FourthDay"
	"AdventOfCode2018/NineteenthDay"
	"AdventOfCode2018/SecondDay"
	"AdventOfCode2018/SeventhDay"
	"AdventOfCode2018/SixteenthDay"
	"AdventOfCode2018/TenthDay"
	"AdventOfCode2018/ThirdDay"
	"AdventOfCode2018/ThirteenthDay"
	"AdventOfCode2018/TwentyFirstDay"
	"fmt"
)

func main() {
	day := 21

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
		fmt.Println(FourthDay.SleepTimeStrategy2())
	case 5:
		fmt.Println(FifthDay.Reduction())
		fmt.Println(FifthDay.ShortestPolymer())
	case 6:
		fmt.Println(EightDay.MetadataSum())
	case 7:
		fmt.Println(SeventhDay.Order())
		fmt.Println(SeventhDay.Duration())
	case 8:
		fmt.Println(EightDay.MetadataSum())
		fmt.Println(EightDay.RootNodeEntry())
	case 10:
		TenthDay.Message()
	case 11:
		fmt.Println(EleventhDay.Coordinates())
		fmt.Println(EleventhDay.CoordinatesAndSize())
	case 13:
		ThirteenthDay.Crash()
		ThirteenthDay.LastOneAlive()
	case 14:
		FourteenthDay.Score()
		FourteenthDay.OtherInterpretation()
	case 16:
		SixteenthDay.SameOpcodes()
		SixteenthDay.ExecuteProgram()
	case 19:
		NineteenthDay.Jumps()
		NineteenthDay.JumpsFrom1()
	case 21:
		TwentyFirstDay.LowestInteger()
	}
}
