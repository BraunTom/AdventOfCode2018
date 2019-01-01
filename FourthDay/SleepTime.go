package FourthDay

import (
	"fmt"
	"io/ioutil"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("FourthDay/input.txt")
	return strings.Split(string(input), "\n")
}

func split(input string) Entry {
	e := Entry{}

	descriptionRegex := regexp.MustCompile("] .*")
	e.description = descriptionRegex.FindString(input)[2:]

	dateRegex := regexp.MustCompile("\\[.*]")
	date := dateRegex.FindString(input)
	date = date[1 : len(date)-1]

	timeRegex := regexp.MustCompile("\\d{2}:\\d{2}")
	time := timeRegex.FindString(date)

	splittedTime := strings.Split(time, ":")

	e.hour, _ = strconv.Atoi(splittedTime[0])
	e.minute, _ = strconv.Atoi(splittedTime[1])

	datumRegex := regexp.MustCompile("-\\d{2}-\\d{2}")
	datum := datumRegex.FindString(date)[1:]
	monthAndDay := strings.Split(datum, "-")

	e.month, _ = strconv.Atoi(monthAndDay[0])
	e.day, _ = strconv.Atoi(monthAndDay[1])

	return e
}

func SleepTime() int {
	var sortedFileData Entries

	for _, v := range fileData() {
		sortedFileData = append(sortedFileData, split(v))
	}

	sort.Sort(sortedFileData)

	for _, v := range sortedFileData {
		fmt.Println(v)
	}

	return 0
}
