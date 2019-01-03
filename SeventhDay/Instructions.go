package SeventhDay

import (
	"fmt"
	"io/ioutil"
	"sort"
	"strings"
)

type workItem struct {
	time uint8
	item *node
}

type workItems []*workItem

func (wI workItems) print() {
	fmt.Print("{")
	for _, v := range wI {
		fmt.Print(v.item.value, ":", v.time, " ")
	}
	fmt.Print("}\n")

}

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

func insertSort(data workItems, el *workItem) workItems {
	index := sort.Search(len(data), func(i int) bool { return data[i].time > el.time })
	data = append(data, &workItem{})
	copy(data[index+1:], data[index:])
	data[index] = el
	return data
}

func (wI workItems) contains(item *workItem) bool {
	for _, v := range wI {
		if v.item.value == item.item.value {
			return true
		}
	}

	return false
}

func Duration() int {
	order := makeGraph()

	for _, v := range fileData() {
		order.addConnection(string(v[5]), string(v[36]))
	}
	order.sort()

	maxWorkerCount := 5
	var inProgress workItems
	time := 0

	for order.size() > 0 {
		sources := order.sources()

		for _, v := range sources {
			if maxWorkerCount-len(inProgress) > 0 {
				item := workItem{v.value[0] + 61 - 'A', v}
				if !inProgress.contains(&item) {
					inProgress = insertSort(inProgress, &item)
				}
			} else {
				break
			}
		}

		nextItem := inProgress[0]

		var nextItemTime = nextItem.time
		time += int(nextItemTime)
		inProgress = inProgress[1:]

		for _, v := range inProgress {
			v.time -= nextItemTime
		}

		order.removeNode(nextItem.item)
	}

	return time
}
