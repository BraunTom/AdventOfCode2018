package ThirteenthDay

type point struct {
	x int
	y int
}

func (p *point) less(otherPoint point) bool {
	if p.y < otherPoint.y {
		return true
	}
	if p.y == otherPoint.y {
		if p.x < otherPoint.x {
			return true
		}
	}

	return false
}

func (p *point) equals(otherPoint point) bool {
	return p.x == otherPoint.x && p.y == otherPoint.y
}
