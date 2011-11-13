package main

import (
	"fmt"
	"math"
)

type color struct {
	r, g, b byte
}

func (me *color) toUint32() uint32 {
	return (uint32(me.r) << 24) | (uint32(me.g) << 16) | (uint32(me.b) << 8)
}

func encode(val uint, bits int) string {
	str := ""
	for i := 0; i < bits; i++ {
		r := val & uint(math.Pow(2, float64(i)))
		if r == 0 {
			str = "0" + str
		} else {
			str = "1" + str
		}
	}
	return str
}

type node struct {
	left  uint8
	color uint32
	right uint8
}

func (me *node) toBinary() string {
	str := encode(uint(me.left), 8)
	str += encode(uint(me.color), 32)
	str += encode(uint(me.right), 8)
	return str
}

func main() {
	tree := make([]node, 1)
	var colors []color
	colors = append(colors, color{255, 0, 0})
	colors = append(colors, color{0, 255, 0})
	colors = append(colors, color{0, 0, 255})
	colors = append(colors, color{255, 255, 0})
	colors = append(colors, color{0, 255, 255})
	colors = append(colors, color{255, 0, 255})
	colors = append(colors, color{255, 255, 255})

	for i := 0; i < len(colors); i++ {
		tree = addNode(0, tree, colors[i])
	}

	for i := 1; i < len(tree); i++ {
		fmt.Print(tree[i].toBinary())
	}
	fmt.Println()
}

func traverse(tree []node, pt uint8) {
	if tree[pt].left != 0 {
		traverse(tree, tree[pt].left)
	}
	fmt.Println("\t", tree[pt].color)
	if tree[pt].right != 0 {
		traverse(tree, tree[pt].right)
	}
}

func addNode(nodePt uint8, tree []node, color color) ([]node) {
	if tree[nodePt].color > color.toUint32() {
		if tree[nodePt].left != 0 {
			tree = addNode(tree[nodePt].left, tree, color)
		} else {
			var node node
			node.color = color.toUint32()
			tree[nodePt].left = uint8(len(tree))
			tree = append(tree, node)
		}
		return tree
	} else {
		if tree[nodePt].right != 0 {
			tree = addNode(tree[nodePt].right, tree, color)
		} else {
			var node node
			node.color = color.toUint32()
			tree[nodePt].right = uint8(len(tree))
			tree = append(tree, node)
		}
		return tree
	}
	return nil
}
