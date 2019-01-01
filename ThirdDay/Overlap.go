package ThirdDay

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("ThirdDay/input.txt")
	return strings.Split(string(input), "\n")
}

func split(line string) (id, x, y, width, height int) {
	splittedLine := strings.Split(line, " @ ")
	id, _ = strconv.Atoi(splittedLine[0][1:])

	splittedRest := strings.Split(splittedLine[1], ": ")

	coordinates := strings.Split(splittedRest[0], ",")
	x, _ = strconv.Atoi(coordinates[0])
	y, _ = strconv.Atoi(coordinates[1])

	extent := strings.Split(splittedRest[1], "x")
	width, _ = strconv.Atoi(extent[0])
	height, _ = strconv.Atoi(extent[1])

	return
}

func Overlap() int {
	var suitArea [1000][1000]int
	overlap := 0

	for _, v := range fileData() {
		_, x, y, width, height := split(v)

		for i := x; i < x+width; i++ {
			for k := y; k < y+height; k++ {
				if suitArea[k][i] == 0 {
					suitArea[k][i] = 1
				} else {
					if suitArea[k][i] == 1 {
						suitArea[k][i] = 2
						overlap++
					}
				}
			}
		}
	}

	return overlap
}

func NoOverlap() int {
	var suitArea [1000][1000]int
	overlappingClaimIDs := make(map[int]bool)

	for _, v := range fileData() {
		id, x, y, width, height := split(v)

		for i := x; i < x+width; i++ {
			for k := y; k < y+height; k++ {
				if suitArea[k][i] == 0 {
					suitArea[k][i] = id
					overlappingClaimIDs[id] = overlappingClaimIDs[id]
				} else {
					overlappingClaimIDs[id] = true
					overlappingClaimIDs[suitArea[k][i]] = true
				}
			}
		}
	}

	for k, v := range overlappingClaimIDs {
		if !v {
			return k
		}
	}

	return 0
}
