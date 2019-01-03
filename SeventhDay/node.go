package SeventhDay

import (
	"fmt"
	"sort"
)

type node struct {
	value    string
	children []*node
	parents  []*node
}

func (n node) hasParent() bool {
	return len(n.parents) > 0
}

func (n node) hasChild() bool {
	return len(n.children) > 0
}

func contains(n *node, nodes *[]*node) bool {
	for _, v := range *nodes {
		if v == n {
			return true
		}
	}
	return false
}

func indexOf(n *node, nodes *[]*node) int {
	for k, v := range *nodes {
		if v == n {
			return k
		}
	}

	return -1
}

func (n *node) addConnectionTo(otherNode *node) {
	if !contains(otherNode, &n.children) {
		n.children = append(n.children, otherNode)
	}
	if !contains(n, &otherNode.parents) {
		otherNode.parents = append(otherNode.parents, n)
	}
}

func (n *node) addConnectionFrom(otherNode *node) {
	otherNode.addConnectionTo(n)
}

func (n *node) removeConnectionTo(otherNode *node) {
	if contains(n, &otherNode.parents) {
		index := indexOf(n, &otherNode.parents)
		otherNode.parents = append(otherNode.parents[:index], otherNode.parents[index+1:]...)
	}
}

func (n *node) removeConnectionFrom(otherNode *node) {
	otherNode.removeConnectionTo(n)
}

func (n *node) debugInfo() string {
	info := ""
	for _, v := range n.children {
		info += "\t" + n.value + " -> " + v.value + ";\n"
	}

	return info
}

func sortNodeSlice(s *[]*node) {
	sort.Slice(*s, func(i, j int) bool {
		return (*s)[i].value < (*s)[j].value
	})
}

func (n *node) sort() {
	sortNodeSlice(&n.children)
	sortNodeSlice(&n.parents)
}

func (n *node) reverse() {
	n.children, n.parents = n.parents, n.children
}

func (n *node) delete() {
	for _, v := range n.children {
		v.removeConnectionFrom(n)
	}
	n.children = []*node{}
	for _, v := range n.parents {
		v.removeConnectionTo(n)
	}
}

func (n *node) print() {
	fmt.Print("{" + n.value + " children: {")
	for _, v := range n.children {
		fmt.Print(v.value + ", ")
	}
	fmt.Print("} parents: {")
	for _, v := range n.parents {
		fmt.Print(v.value + ", ")
	}
	fmt.Print("}}\n")
}
