package SixteenthDay

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("SixteenthDay/input.txt")
	return strings.Split(string(input), "\n")
}

func SameOpcodes() {
	data := fileData()[:3127]
	var samples [][3][4]int

	for i := 0; i < len(data); i += 4 {
		var entry [3][4]int

		var before [4]int

		for k, v := range strings.Split(data[i][9:19], ", ") {
			number, _ := strconv.Atoi(v)
			before[k] = number
		}

		var instruction [4]int
		for k, v := range strings.Split(data[i+1], " ") {
			number, _ := strconv.Atoi(v)
			instruction[k] = number
		}

		var after [4]int
		for k, v := range strings.Split(data[i+2][9:19], ", ") {
			number, _ := strconv.Atoi(v)
			after[k] = number
		}

		entry[0] = before
		entry[1] = instruction
		entry[2] = after

		samples = append(samples, entry)
	}

	sum := 0
loop:
	for _, v := range samples {
		count := 0
		for e := range tests {
			if tests[e](v[0], v[1], v[2]) {
				count++
			}
			if count >= 3 {
				sum++
				continue loop
			}
		}
	}

	fmt.Println(sum)
}

func intersection(slice1, slice2 []int) []int {
	doesExist := make(map[int]bool)
	var i []int

	for _, v := range slice1 {
		doesExist[v] = true
	}

	for _, v := range slice2 {
		if doesExist[v] {
			i = append(i, v)
		}
	}

	return i
}

func opcodeToProcedure() {
	data := fileData()[:3127]
	var samples [][3][4]int

	for i := 0; i < len(data); i += 4 {
		var entry [3][4]int

		var before [4]int

		for k, v := range strings.Split(data[i][9:19], ", ") {
			number, _ := strconv.Atoi(v)
			before[k] = number
		}

		var instruction [4]int
		for k, v := range strings.Split(data[i+1], " ") {
			number, _ := strconv.Atoi(v)
			instruction[k] = number
		}

		var after [4]int
		for k, v := range strings.Split(data[i+2][9:19], ", ") {
			number, _ := strconv.Atoi(v)
			after[k] = number
		}

		entry[0] = before
		entry[1] = instruction
		entry[2] = after

		samples = append(samples, entry)
	}

	m := make(map[int][]int)
	for _, v := range samples {
		count := 0
		var possibilities []int
		for e := range tests {
			if tests[e](v[0], v[1], v[2]) {
				count++
				possibilities = append(possibilities, e)
			}
		}
		if m[v[1][0]] == nil {
			m[v[1][0]] = []int{}
		}
		if len(m[v[1][0]]) == 0 {

			m[v[1][0]] = possibilities
		}
		if len(m[v[1][0]]) == 1 {
			m[v[1][0]] = intersection(m[v[1][0]], possibilities)
		}

	}

	for k, v := range m {
		fmt.Println(k, v, "\n")
	}
}

func ExecuteProgram() {
	data := fileData()[3129:]
	var m machine

	for _, v := range data {
		var instruction [4]int
		for k, v := range strings.Split(v, " ") {
			number, _ := strconv.Atoi(v)
			instruction[k] = number
		}

		m.executeInstruction(instruction)
	}

	fmt.Println(m[0])
}
