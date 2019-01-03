package SeventhDay

import (
	"io/ioutil"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("SeventhDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Order() string {
	order := makeGraph()

	for _, v := range fileData() {
		order.addConnection(string(v[5]), string(v[36]))
	}
	order.sort()

	var solution []string

	for order.size() > 0 {
		start := order.sources()[0]
		solution = append(solution, start.value)
		order.removeNode(start)
	}

	return strings.Join(solution, "")
}

func Duration() {
	order := makeGraph()

	for _, v := range fileData() {
		order.addConnection(string(v[5]), string(v[36]))
	}
	order.sort()

}
