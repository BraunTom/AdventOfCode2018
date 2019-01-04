package EleventhDay

type summedAreaTable [height][width]int

func (sAT *summedAreaTable) init(grid [height][width]int) {
	for y := 0; y < height; y++ {
		sum := 0
		for x := 0; x < width; x++ {
			if y == 0 {
				sAT[y][x] = sum + grid[y][x]
			} else {
				sAT[y][x] = sum + grid[y][x] + sAT[y-1][x]
			}
			sum += grid[y][x]
		}
	}
}

func (sAT *summedAreaTable) at(x, y int) int {
	if x < 0 || y < 0 || x >= width || y >= height {
		return 0
	} else {
		return sAT[y][x]
	}
}

func (sAT *summedAreaTable) getArea(x, y, width, height int) int {
	x -= 1
	y -= 1
	return sAT.at(x+width, y+height) - sAT.at(x+width, y) - sAT.at(x, y+height) + sAT.at(x, y)
}
