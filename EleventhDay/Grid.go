package EleventhDay

import (
	"math"
)

const (
	input  = 7347
	width  = 300
	height = 300
)

func Coordinates() [2]int {
	var grid [height][width]int
	var summedAreas summedAreaTable

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rackId := x + 11
			powerLevel := (rackId*(y+1) + input) * rackId
			grid[y][x] = ((powerLevel / 100) % 10) - 5
		}
	}

	summedAreas.init(grid)

	var max = math.MinInt64
	var coordinates [2]int

	for y := 1; y < height-1; y++ {
		for x := 1; x < width-1; x++ {
			value := summedAreas.getArea(x-1, y-1, 3, 3)

			if value > max {
				max, coordinates = value, [2]int{x, y}
			}
		}
	}

	return coordinates
}

func CoordinatesAndSize() [3]int {
	var grid [height][width]int
	var summedAreas summedAreaTable

	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			rackId := x + 11
			powerLevel := (rackId*(y+1) + input) * rackId
			grid[y][x] = ((powerLevel / 100) % 10) - 5
		}
	}

	summedAreas.init(grid)

	var max = math.MinInt64
	var coordinates [3]int

	for size := 1; size <= 300; size++ {
		for y := size / 2; y < height-((size+1)/2); y++ {
			for x := size / 2; x < width-((size+1)/2); x++ {

				value := summedAreas.getArea(x-((size+1)/2), y-((size+1)/2), size, size)

				if value > max {
					max = value
					coordinates = [3]int{x + 1 - size/2, y + 1 - size/2, size}
				}
			}
		}

	}

	return coordinates
}
