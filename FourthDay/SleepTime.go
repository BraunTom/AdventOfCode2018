package FourthDay

import (
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

func split(input string) entry {
	e := entry{}

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
	var sortedFileData entries
	var timeline [][60]int
	minutesAsleep := make(map[int]int)

	for _, v := range fileData() {
		sortedFileData = append(sortedFileData, split(v))
	}

	sort.Sort(sortedFileData)

	guardNumber := sortedFileData[0].ID()
	var currentRow [60]int
	for i := 0; i < len(sortedFileData); i++ {
		if sortedFileData[i].Type() == shift {
			guardNumber = sortedFileData[i].ID()
			if i != 0 {
				timeline = append(timeline, currentRow)
			}
			currentRow = [60]int{}
		}
		if sortedFileData[i].Type() == asleep {
			for k := sortedFileData[i].minute; k < sortedFileData[i+1].minute; k++ {
				minutesAsleep[guardNumber]++
				currentRow[k] = guardNumber
			}
			i++
		}
	}
	/*
		for i := 0; i < 60; i++ {
			fmt.Printf("%5d", i)
		}
		fmt.Print("\n")
		for _, v := range timeline {
			for _, v := range v {
				fmt.Printf("%5d", v)
			}
			fmt.Print("\n")
		}*/

	maxKey := 0
	maxValue := 0
	for k, v := range minutesAsleep {
		if v > maxValue {
			maxValue = v
			maxKey = k
		}
	}

	maxColumnValue := 0
	maxColumnIndex := 0
	for i := 0; i < 60; i++ {
		count := 0
		for k := 0; k < len(timeline); k++ {
			if timeline[k][i] == maxKey {
				count++
			}
		}

		if count > maxColumnValue {
			maxColumnValue = count
			maxColumnIndex = i
		}

	}

	return maxKey * maxColumnIndex
}

func SleepTimeStrategy2() int {
	var sortedFileData entries
	var timeline [][60]int
	minutesAsleep := make(map[int]*[60]int) // guardID => minute => count

	for _, v := range fileData() {
		sortedFileData = append(sortedFileData, split(v))
	}

	sort.Sort(sortedFileData)

	maxID := 0
	maxValue := 0
	maxTime := 0

	guardNumber := sortedFileData[0].ID()
	var currentRow [60]int
	for i := 0; i < len(sortedFileData); i++ {
		if sortedFileData[i].Type() == shift {
			guardNumber = sortedFileData[i].ID()
			if i != 0 {
				timeline = append(timeline, currentRow)
			}
			currentRow = [60]int{}
		}
		if sortedFileData[i].Type() == asleep {
			for k := sortedFileData[i].minute; k < sortedFileData[i+1].minute; k++ {
				if minutesAsleep[guardNumber] == nil {
					minutesAsleep[guardNumber] = &[60]int{}
				}
				minutesAsleep[guardNumber][k]++

				if minutesAsleep[guardNumber][k] > maxValue {
					maxValue = minutesAsleep[guardNumber][k]
					maxID = guardNumber
					maxTime = k
				}

				currentRow[k] = guardNumber
			}
			i++
		}
	}

	return maxID * maxTime
}
