package FirstDay

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("FirstDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Frequency() int {
	sum := 0

	for _, v := range fileData() {
		number, _ := strconv.Atoi(v)
		sum += number
		fmt.Println(sum)
	}

	return sum
}

func Twice() int {
	set := make(map[int]bool)
	data := fileData()
	sum := 0

	set[0] = true

	for {
		for _, v := range data {
			number, err := strconv.Atoi(v)

			if err != nil {
				continue
			}

			sum += number

			if set[sum] {
				return sum
			}

			set[sum] = true
		}
	}
}
