package SeventhDay

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"runtime"
)

type graph struct {
	nodes   []*node
	mapping map[string]*node
}

func makeGraph() graph {
	return graph{mapping: make(map[string]*node)}
}

func (g *graph) size() int {
	return len(g.nodes)
}

func (g *graph) sources() []*node {
	var nodes []*node
	for _, v := range g.nodes {
		if !v.hasParent() {
			nodes = append(nodes, v)
		}
	}

	return nodes
}

func (g *graph) addNode(n *node) {
	g.nodes = append(g.nodes, n)
	g.mapping[n.value] = n
}

func (g *graph) removeID(id string) *node {
	return g.removeNode(g.getNode(id))
}

func (g *graph) removeNode(n *node) *node {
	nodeIndex := indexOf(n, &g.nodes)

	if nodeIndex == -1 {
		return nil
	}

	g.nodes = append(g.nodes[:nodeIndex], g.nodes[nodeIndex+1:]...)
	n.delete()

	return n
}

func (g graph) getNode(id string) *node {
	return g.mapping[id]
}

func (g *graph) addConnection(from, to string) {
	if g.mapping[from] == nil {
		g.addNode(&node{value: from})
	}

	if g.mapping[to] == nil {
		g.addNode(&node{value: to})
	}

	g.getNode(from).addConnectionTo(g.getNode(to))
}

func (g *graph) reverse() {
	for _, v := range g.nodes {
		v.reverse()
	}
}

func (g *graph) sort() {
	sortNodeSlice(&g.nodes)

	for _, v := range g.nodes {
		v.sort()
	}
}

func (g *graph) print() {
	for _, v := range g.nodes {
		v.print()
	}
}

func (g *graph) debug(filename string) {
	dotFilename := filename + ".dot"

	info := "digraph{\n"
	for _, v := range g.nodes {
		info += v.debugInfo()
	}
	info += "}"
	err := ioutil.WriteFile("SeventhDay/"+dotFilename, []byte(info), 0644)

	if err != nil {
		panic(err)
	}

	_, b, _, _ := runtime.Caller(0)
	basepath := filepath.Dir(b)

	fmt.Println(basepath)

	cmd := exec.Command("dot", "-Tpng", basepath+"/"+dotFilename, "-o", basepath+"/"+filename+".png")

	err = cmd.Run()
	if err != nil {
		panic(err)
	}
}
