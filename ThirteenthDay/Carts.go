package ThirteenthDay

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("ThirteenthDay/input.txt")
	return strings.Split(string(input), "\n")
}

func Crash() {
	var simulation cartSimulation
	simulation.init(fileData())

	cartCount := len(simulation.carts)

	for cartCount == len(simulation.carts) {
		simulation.step()

		/*
			simulation.print()

			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')//*/
	}
}

func LastOneAlive() {
	var simulation cartSimulation
	simulation.init(fileData())

	for len(simulation.carts) > 1 {
		simulation.step()

		/*
			simulation.print()
			fmt.Println("//////////////////////////")*/
	}

	fmt.Println(*simulation.carts[0].position)
}
