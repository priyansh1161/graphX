package main

import (
	"image/color"
	"fmt"
)

type Node struct {
	value int
	name string
	weight float32
	color color.RGBA
}

func(n *Node) String() string {
	return "The Value is " + fmt.Sprintf("%d", n.value)
}