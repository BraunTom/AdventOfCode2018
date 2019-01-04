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

func Crash() point {
	var simulation cartSimulation
	simulation.init(fileData())

	for !simulation.hasCrash() {
		simulation.step()

		/*
			simulation.print()

			reader := bufio.NewReader(os.Stdin)
			reader.ReadString('\n')//*/
	}

	return simulation.getCrash()
}

func LastOneAlive() point {
	var simulation cartSimulation
	simulation.init(fileData())

	i := 0
	for len(simulation.carts) > 1 {
		simulation.step()

		if simulation.hasCrash() {
			fmt.Println(i)
			simulation.removeCrashed()
		}

		i++
	}

	simulation.step()

	return *simulation.carts[0].position
}
