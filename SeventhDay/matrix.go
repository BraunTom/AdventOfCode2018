package SeventhDay

import (
	"fmt"
)

// implemented after misunderstanding the task first

type matrix [][]int

func (ma *matrix) initialize(n, m int) {
	for i := 0; i < n; i++ {
		var row []int
		for k := 0; k < m; k++ {
			row = append(row, 0)
		}

		*ma = append(*ma, row)
	}
}

func (ma *matrix) deleteRow(index int) {
	*ma = append((*ma)[:index], (*ma)[index+1:]...)
}

func (ma *matrix) deleteColumn(index int) {
	for k, v := range *ma {
		(*ma)[k] = append(v[:index], v[index+1:]...)
	}
}

func (ma matrix) column(index int) []int {
	var column []int

	for _, v := range ma {
		column = append(column, v[index])
	}

	return column
}

func (ma matrix) row(index int) []int {
	return ma[index]
}

func (ma *matrix) fillRowWith(index, number int) {
	for i := 0; i < ma.width(); i++ {
		(*ma)[index][i] = number
	}
}

func (ma *matrix) fillColumnWith(index, number int) {
	for i := 0; i < ma.height(); i++ {
		(*ma)[i][index] = number
	}
}

func (ma matrix) isRowFilledWith(index, number int) bool {
	for _, v := range ma.row(index) {
		if v != number {
			return false
		}
	}

	return true
}

func (ma matrix) isColumnFilledWith(index, number int) bool {
	for _, v := range ma.column(index) {
		if v != number {
			return false
		}
	}

	return true
}

func (ma matrix) columnContainsOnly(index int, allowed []int) bool {
loop:
	for _, v := range ma.column(index) {
		for _, number := range allowed {
			if v == number {
				continue loop
			}
		}
		return false
	}

	return true
}

func (ma matrix) print() {
	fmt.Print("   ")
	for i := 0; i < 26; i++ {
		fmt.Print(string(i+'A'), " ")

	}

	fmt.Print("\n")

	for k, v := range ma {
		fmt.Println(string(k+'A'), v)
	}
}

func (ma matrix) width() int {
	return len(ma[0])
}

func (ma matrix) height() int {
	return len(ma)
}
