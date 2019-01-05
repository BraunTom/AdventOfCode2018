package TwentyFirstDay

import (
	"io/ioutil"
	"strconv"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("TwentyFirstDay/input.txt")
	return strings.Split(string(input), "\n")
}

/*
	do {
		b = 123
		b = b & 456
	} while (b != 72)

	b = 0
	do {
		e = b | 65.536
		while (true){
			b = 8.595.037
			d = e & 255		<= d = e % 256
			b += d
			b = b & 16.777.215
			b *= 65.899
			b = b & 16.777.215
			if(e >= 256) {
				for (d = 0; c <= e; d++){
					c = d + 1
					c *= 256
				} 		<= e = floor(e / 256)
			} else {
				break
			}
			e = d
		}

	} while ( b != a)
*/

func LowestInteger() {

	var m machine
	data := fileData()
	var instructions []instruction

	index, _ := strconv.Atoi(string(data[0][4]))
	m.bindIp(index)
	// m.register[0] = 15883666

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
}
