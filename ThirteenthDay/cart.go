package ThirteenthDay

import "fmt"

type turnState int
type direction int

const (
	turnLeft turnState = iota
	goStreight
	goRight
)

const (
	left direction = iota
	right
	up
	down
)

type cart struct {
	position *point

	dir direction
	tS  turnState
}

func isInputCart(input string) bool {
	switch input {
	case "<", ">", "v", "^":
		return true
	default:
		return false
	}
}

func makeCartFrom(x, y int, input string) *cart {
	var direction direction

	switch input {
	case "<":
		direction = left
	case ">":
		direction = right
	case "v":
		direction = down
	case "^":
		direction = up
	}

	return &cart{&point{x, y}, direction, turnLeft}
}

func (c *cart) print() {
	fmt.Println("{", c.position.x, ",", c.position.y, "}", c.dir)
}

func (c *cart) moveOn(t *tile) {
	switch t.direction {
	case straightHorizontal:
		switch c.dir {
		case left:
			c.position.x -= 1
		case right:
			c.position.x += 1
		default:
			panic("wtf")
		}

	case straightVertical:
		switch c.dir {
		case up:
			c.position.y -= 1
		case down:
			c.position.y += 1
		default:
			panic("wtf")
		}
	case curveLeft:
		switch c.dir {
		case up:
			c.dir = left
			c.position.x -= 1
		case down:
			c.dir = right
			c.position.x += 1
		case left:
			c.dir = up
			c.position.y -= 1
		case right:
			c.dir = down
			c.position.y += 1
		}
	case curveRight:
		switch c.dir {
		case up:
			c.dir = right
			c.position.x += 1
		case down:
			c.dir = left
			c.position.x -= 1
		case left:
			c.dir = down
			c.position.y += 1
		case right:
			c.dir = up
			c.position.y -= 1
		}
	case intersection:
		switch c.dir {
		case up:
			switch c.tS {
			case turnLeft:
				c.dir = left
				c.position.x -= 1
			case goStreight:
				c.position.y -= 1
			case goRight:
				c.dir = right
				c.position.x += 1
			}
		case down:
			switch c.tS {
			case turnLeft:
				c.dir = right
				c.position.x += 1
			case goStreight:
				c.position.y++
			case goRight:
				c.dir = left
				c.position.x -= 1
			}
		case left:
			switch c.tS {
			case turnLeft:
				c.dir = down
				c.position.y += 1
			case goStreight:
				c.position.x--
			case goRight:
				c.dir = up
				c.position.y -= 1
			}
		case right:
			switch c.tS {
			case turnLeft:
				c.dir = up
				c.position.y -= 1
			case goStreight:
				c.position.x++
			case goRight:
				c.dir = down
				c.position.y += 1
			}
		}
		c.tS = (c.tS + 1) % 3
	}
}
