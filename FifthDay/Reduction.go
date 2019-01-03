package FifthDay

import (
	"io/ioutil"
	"math"
	"strings"
)

func fileData() string {
	input, _ := ioutil.ReadFile("FifthDay/input.txt")
	return string(input)
}

func doReact(a, b uint8) bool {
	return math.Abs(float64(a)-float64(b)) == 32
}

func reduce(input string) string {
	var reducedData []string
	dataLen := len(input)
	reacted := false

	for i := 0; i < dataLen; i++ {
		if reacted && len(reducedData) > 0 && doReact(reducedData[len(reducedData)-1][0], input[i]) {
			reducedData = reducedData[:len(reducedData)-1]
		} else {
			if i != dataLen-1 && doReact(input[i], input[i+1]) {
				i++
				reacted = true
			} else {
				reducedData = append(reducedData, string(input[i]))
				reacted = false
			}
		}

	}

	return strings.Join(reducedData, "")
}

func Reduction() int {
	return len(reduce(fileData()))
}

func ShortestPolymer() int {
	data := fileData()
	min := len(data)
	for i := 0; i < 26; i++ {
		shortenedPolymer := strings.Replace(data, string('a'+i), "", -1)
		shortenedPolymer = strings.Replace(shortenedPolymer, string('A'+i), "", -1)

		l := len(reduce(shortenedPolymer))

		if l < min {
			min = l
		}
	}

	return min
}
