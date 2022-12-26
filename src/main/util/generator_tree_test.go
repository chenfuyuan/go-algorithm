package util

import (
	"fmt"
	"testing"
)

func TestGeneratorTree(t *testing.T) {
	source := []int{5, 1, 4, -1, -1, 3, 6}
	tree := GeneratorTree(source)
	fmt.Print(tree)
}
