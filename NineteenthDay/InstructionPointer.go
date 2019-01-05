package NineteenthDay

import (
	"fmt"
	"io/ioutil"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("NineteenthDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Jumps() {
	var m machine
	data := fileData()
	var instructions []instruction

	index, _ := strconv.Atoi(string(data[0][4]))
	m.bindIp(index)

	for _, v := range data[1:] {
		var ins instruction
		split := strings.Split(v, " ")

		ins[0] = instructionToOpcode(split[0])

		for i := 1; i < len(split); i++ {
			number, _ := strconv.Atoi(split[i])
			ins[i] = number
		}

		instructions = append(instructions, ins)
	}

	m.execute(instructions)
	fmt.Println(m.register[0])
}

func JumpsFrom1() {
	var m machine
	data := fileData()
	var instructions []instruction

	index, _ := strconv.Atoi(string(data[0][4]))
	m.bindIp(index)
	m.register[0] = 1

	for _, v := range data[1:] {
		var ins instruction
		split := strings.Split(v, " ")

		ins[0] = instructionToOpcode(split[0])

		for i := 1; i < len(split); i++ {
			number, _ := strconv.Atoi(split[i])
			ins[i] = number
		}

		instructions = append(instructions, ins)
	}

	// m.execute(instructions) to slow
	/*
		the program is calculating the sum of all deviders of 10551394 in a very inefficient way
	*/
	// 10551394 = 2 * 7 * 167 * 4513
	fmt.Println(1 + 2 + 7 + 14 + 167 + 334 + 1169 + 2338 + 4513 + 9026 + 31591 + 63182 + 753671 + 1507342 + 5275697 + 10551394)
}
