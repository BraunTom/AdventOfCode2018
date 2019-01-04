package EightDay

import (
	"io/ioutil"
	"strings"
)

func fileData() []string {
	input, _ := ioutil.ReadFile("EightDay/input.txt")
	return strings.Split(string(input), " ")
}

func MetadataSum() int {
	return makeNode(fileData()).metadataSum()
}

func RootNodeEntry() int {
	return makeNode(fileData()).value()
}
