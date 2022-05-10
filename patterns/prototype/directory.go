package main

import "fmt"

type directory struct {
	children []inode
	name     string
}

func (d *directory) print(indentation string) {
	fmt.Println(indentation + d.name)

	for _, i := range d.children {
		i.print(indentation + indentation)
	}
}

func (d *directory) clone() inode {
	cloneDir := &directory{name: d.name + "_clone"}

	var tmpChildren []inode
	for _, i := range d.children {
		clone := i.clone()
		tmpChildren = append(tmpChildren, clone)
	}

	cloneDir.children = tmpChildren
	return cloneDir
}
