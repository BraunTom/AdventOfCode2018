package ThirdDay

import (
	"io/ioutil"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("ThirdDay/input.txt")
	return strings.Split(string(input), "\n")
}

/*
func split(line string) (id, x, y, width, height int){
	splittedLine := strings.Split(line, " @ ")
	id, _ = strconv.Atoi(splittedLine[0])

	strings.Split(splittedLine[1], ": ")
}*/

func Overlap() int {
	//square := [1000][1000]int

	return 0
}
