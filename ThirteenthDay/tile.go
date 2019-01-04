package ThirteenthDay

type tileDirection int

const (
	straightHorizontal tileDirection = iota
	straightVertical
	curveLeft
	curveRight
	intersection
)

type tile struct {
	position  point
	direction tileDirection
}

func isInputTile(input string) bool {
	switch input {
	case "|", "-", "/", "\\", "+":
		return true
	default:
		return false
	}
}

func makeTileFrom(x, y int, input string) *tile {
	var direction tileDirection

	switch input {
	case "|", "v", "^":
		direction = straightVertical
	case "-", "<", ">":
		direction = straightHorizontal
	case "/":
		direction = curveRight
	case "\\":
		direction = curveLeft
	case "+":
		direction = intersection
	}

	return &tile{point{x, y}, direction}
}
