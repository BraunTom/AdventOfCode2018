package FourteenthDay

import (
	"fmt"
	"math"
)

const (
	input        = 37
	after        = 147061
	recipeNumber = 10
)

func equals(slice1, slice2 []int) bool {
	for k, v := range slice1 {
		if v != slice2[k] {
			return false
		}
	}

	return true
}

func numberLen(number int) int {
	return int(math.Log10(float64(number)) + 1)
}

func toSlice(number int) []int {
	slice := make([]int, numberLen(number))

	for i := len(slice) - 1; i >= 0; i-- {
		slice[i] = number % 10
		number /= 10
	}

	return slice
}

func Score() {
	scoreBoard := toSlice(input)

	fmt.Println(scoreBoard)

	elve1Position := 0
	elve2Position := 1

	for len(scoreBoard) < after+recipeNumber {
		newScore := scoreBoard[elve1Position] + scoreBoard[elve2Position]

		if newScore < 10 {
			scoreBoard = append(scoreBoard, newScore)
		} else {
			scoreBoard = append(scoreBoard, newScore/10, newScore%10)
		}

		elve1Position = (elve1Position + scoreBoard[elve1Position] + 1) % len(scoreBoard)
		elve2Position = (elve2Position + scoreBoard[elve2Position] + 1) % len(scoreBoard)

		//fmt.Println(scoreBoard)
	}

	fmt.Println(scoreBoard[after : after+recipeNumber])
}

func OtherInterpretation() {
	scoreBoard := toSlice(input)
	searchedValue := toSlice(after)
	searchedLen := numberLen(after)

	fmt.Println(scoreBoard)

	elve1Position := 0
	elve2Position := 1

	for len(scoreBoard) < 50000000 {
		newScore := scoreBoard[elve1Position] + scoreBoard[elve2Position]

		if newScore < 10 {
			scoreBoard = append(scoreBoard, newScore)
			if len(scoreBoard) >= searchedLen && equals(searchedValue, scoreBoard[len(scoreBoard)-searchedLen:]) {
				fmt.Println(scoreBoard[len(scoreBoard)-searchedLen:])
				break
			}
		} else {
			scoreBoard = append(scoreBoard, newScore/10, newScore%10)
			if len(scoreBoard) >= searchedLen && equals(searchedValue, scoreBoard[len(scoreBoard)-searchedLen:]) {
				fmt.Println(len(scoreBoard) - searchedLen)
				break
			}
			if len(scoreBoard)-1 >= searchedLen && equals(searchedValue, scoreBoard[len(scoreBoard)-searchedLen-1:]) {
				fmt.Println(len(scoreBoard) - searchedLen - 1)
				break
			}
		}

		elve1Position = (elve1Position + scoreBoard[elve1Position] + 1) % len(scoreBoard)
		elve2Position = (elve2Position + scoreBoard[elve2Position] + 1) % len(scoreBoard)

		//fmt.Println(scoreBoard)
	}

}
