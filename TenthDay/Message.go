package TenthDay

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("TenthDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Message() {
	data := fileData()
	l := len(data)
	points := make([]*point, l)

	for k, v := range data {
		points[k] = makePoint(string(v))
	}

	for _, v := range points {
		v.simulateSteps(10470)
	}

	i := 0

	for {
		reader := bufio.NewReader(os.Stdin)
		reader.ReadString('\n')

		smallestX := points[0].x
		smallestY := points[0].y
		biggestX := points[0].x
		biggestY := points[0].y

		for _, v := range points {
			if v.x < smallestX {
				smallestX = v.x
			}
			if v.y < smallestY {
				smallestY = v.y
			}

			if v.x > biggestX {
				biggestX = v.x
			}
			if v.y > biggestY {
				biggestY = v.y
			}
		}

		width := biggestX - smallestX
		height := biggestY - smallestY

		canvas := make([]bool, width*(height+1))

		for _, v := range points {
			canvas[(v.x-smallestX)+(v.y-smallestY)*width] = true
		}

		for k, v := range canvas {
			if k%width == 0 {
				fmt.Print("\n")
			}
			if v == true {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}

		}

		fmt.Print("\n\n\n")

		for _, v := range points {
			v.simulateSteps(1)
		}
		fmt.Println(10470 + i)
		i++

	}

}
