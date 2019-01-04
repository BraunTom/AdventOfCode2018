package EightDay

import (
	"fmt"
	"strconv"
)

type node struct {
	childNodeCount int
	metaDataCount  int
	length         int

	children []*node
	metadata []int
}

func makeNode(input []string) *node {
	n := &node{}

	n.childNodeCount, _ = strconv.Atoi(input[0])
	n.metaDataCount, _ = strconv.Atoi(input[1])

	offset := 0
	for i := 0; i < n.childNodeCount; i++ {
		newNode := makeNode(input[2+offset:])
		n.children = append(n.children, newNode)
		offset += newNode.length
	}

	for i := 0; i < n.metaDataCount; i++ {
		var metadataEntry int
		metadataEntry, _ = strconv.Atoi(input[2+n.childLength()+i])
		n.metadata = append(n.metadata, metadataEntry)
	}

	n.length = 2 + n.metaDataCount + n.childLength()
	return n
}

func (n *node) childLength() int {
	length := 0
	for _, v := range n.children {
		length += v.length
	}

	return length
}

func (n *node) print() {
	fmt.Print("{", n.childNodeCount, " ", n.metaDataCount, " children: [")
	for _, v := range n.children {
		fmt.Print("\n")
		v.print()
	}
	fmt.Print("]}\n")
}

func (n *node) metadataSum() int {
	sum := 0
	for _, v := range n.metadata {
		sum += v
	}

	for _, v := range n.children {
		sum += v.metadataSum()
	}

	return sum
}

func (n *node) value() int {
	sum := 0

	if n.childNodeCount == 0 {
		for _, v := range n.metadata {
			sum += v
		}
	} else {
		for _, v := range n.metadata {
			if v <= n.childNodeCount {
				sum += n.children[v-1].value()
			}
		}
	}

	return sum
}

/*
for i := n.metaDataCount; i > 0; i-- {
		var metadataEntry int
		metadataEntry, _ = strconv.Atoi(input[len(input) - i])
		n.metadata = append(n.metadata, metadataEntry)
	}
*/
