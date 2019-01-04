package ThirteenthDay

import (
	"fmt"
	"sort"
)

type cartSimulation struct {
	cartMap [][]*tile
	carts   []*cart
}

func (cM *cartSimulation) init(lines []string) {
	for k, v := range lines {
		v = string(v)
		row := make([]*tile, len(v))

		for e := range v {
			content := string(v[e])
			if isInputTile(content) {
				row[e] = makeTileFrom(e, k, content)
			}
			if isInputCart(content) {
				cM.carts = append(cM.carts, makeCartFrom(e, k, content))
				row[e] = makeTileFrom(e, k, content)
			}
		}
		cM.cartMap = append(cM.cartMap, row)
	}
}

func (cM *cartSimulation) sortCarts() {
	sort.Slice(cM.carts, func(i, j int) bool { return cM.carts[i].position.less(*cM.carts[j].position) })
}

func (cM *cartSimulation) tileAt(p point) *tile {
	return cM.cartMap[p.y][p.x]
}

func (cM *cartSimulation) moveCart(c *cart) {
	c.moveOn(cM.tileAt(*c.position))
}

func (cM *cartSimulation) step() {
	cM.sortCarts()
	var toDelete []*cart
	for _, v := range cM.carts {
		cM.moveCart(v)
		if b, otherCart := cM.hasCrash(v); b {
			fmt.Println("Collision at:", *v.position)
			toDelete = append(toDelete, v, otherCart)
		}
	}

	var deleteIndexes []int
	for _, v := range toDelete {
		for k, c := range cM.carts {
			if c == v {
				deleteIndexes = append(deleteIndexes, k)
			}
		}
	}

	sort.Ints(deleteIndexes)

	for i := len(deleteIndexes) - 1; i >= 0; i-- {
		cM.carts = append(cM.carts[:deleteIndexes[i]], cM.carts[deleteIndexes[i]+1:]...)
	}
}

func (cM *cartSimulation) hasCrash(c *cart) (bool, *cart) {
	for _, v := range cM.carts {
		if c != v && c.position.equals(*v.position) {
			return true, v
		}
	}

	return false, nil
}

func order(x, y int) (int, int) {
	if x > y {
		return x, y
	} else {
		return y, x
	}
}

func (cM *cartSimulation) removeCrashed() {
	for i := 0; i < len(cM.carts); i++ {
		for k := i + 1; k < len(cM.carts); k++ {
			if cM.carts[i].position.equals(*cM.carts[k].position) {
				max, min := order(i, k)
				cM.carts = append(cM.carts[:max], cM.carts[max+1:]...)
				cM.carts = append(cM.carts[:min], cM.carts[min+1:]...)
				return
			}
		}
	}
}

func (cM *cartSimulation) getCrash() point {
	cM.sortCarts()
	for i := 1; i < len(cM.carts); i++ {
		if cM.carts[i-1].position.equals(*cM.carts[i].position) {
			return *cM.carts[i-1].position
		}
	}

	panic("should not happen")
}

func (cM *cartSimulation) print() {
	carts := make(map[int]map[int]*cart)

	for _, v := range cM.carts {
		if carts[v.position.x] == nil {
			carts[v.position.x] = make(map[int]*cart)
		}
		carts[v.position.x][v.position.y] = v
	}

	for _, v := range cM.cartMap {
		for _, v := range v {
			if v != nil {
				if carts[v.position.x][v.position.y] != nil {
					switch carts[v.position.x][v.position.y].dir {
					case left:
						fmt.Print("<")
					case right:
						fmt.Print(">")
					case up:
						fmt.Print("^")
					case down:
						fmt.Print("v")
					}
					continue
				}

				switch v.direction {
				case straightHorizontal:
					fmt.Print("-")
				case straightVertical:
					fmt.Print("|")
				case curveLeft:
					fmt.Print("\\")
				case curveRight:
					fmt.Print("/")
				case intersection:
					fmt.Print("+")
				}
			} else {
				fmt.Print(" ")
			}
		}
		fmt.Print("\n")
	}
}
