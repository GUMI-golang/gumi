package pipelines

import "github.com/GUMI-golang/gumi"

type Node struct {
	Parent   *Node
	Childrun []*Node
	//
	Elem gumi.GUMI
}