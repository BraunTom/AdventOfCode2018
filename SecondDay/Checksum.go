package SecondDay

import (
	"io/ioutil"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("SecondDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Checksum() int {
	twos := 0
	trees := 0

	for _, v := range fileData() {
		charCount := make(map[int32]int)
		for _, v := range v {
			charCount[v]++
		}

		two := false
		tree := false

		for _, v := range charCount {
			if v == 2 && !two {
				twos++
				two = true
			}
			if v == 3 && !tree {
				trees++
				tree = true
			}
		}
	}
	return twos * trees
}

func stringDistance(string1, string2 string) int {
	sum := 0
	for k, v := range string1 {
		if uint8(v) != string2[k] {
			sum++
		}
	}

	return sum
}

func overlappingStringPart(string1, string2 string) string {
	result := ""

	for k, v := range string1 {
		if uint8(v) == string2[k] {
			result += string(v)
		}
	}

	return result
}

func OneApart() string {
	data := fileData()

	for k, v := range data {
		for i := k + 1; i < len(data); i++ {
			if stringDistance(v, data[i]) == 1 {
				return overlappingStringPart(v, data[i])
			}
		}
	}
	return "foo"
}
