package TenthDay

import (
	"strconv"
	"strings"
)

type point struct {
	x  int
	y  int
	dx int
	dy int
}

func makePoint(description string) *point {
	x, _ := strconv.Atoi(strings.TrimSpace(description[10:16]))
	y, _ := strconv.Atoi(strings.TrimSpace(description[18:24]))
	dx, _ := strconv.Atoi(strings.TrimSpace(description[36:38]))
	dy, _ := strconv.Atoi(strings.TrimSpace(description[40:42]))
	return &point{x, y, dx, dy}
}

func (p *point) simulateSteps(n int) {
	p.x += p.dx * n
	p.y += p.dy * n
}
