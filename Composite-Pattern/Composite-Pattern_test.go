package Composite_Pattern

import (
	"container/list"
	"fmt"
	"testing"
)

func TestCompositePattern(t *testing.T) {
	var (
		root   IFile
		readme IFile
		cs     IFile
		csapp  IFile
		apue   IFile
	)
	root = &Folder{
		fileName: "root",
		childs:   list.List{},
	}
	readme = &File{fileName: "README.md"}
	root.Touch(readme)

	cs = &Folder{
		fileName: "Computer_Science",
		childs:   list.List{},
	}
	root.Touch(cs)

	csapp = &File{fileName: "CSAPP.pdf"}
	apue = &File{fileName: "APUE.pdf"}
	cs.Touch(csapp)
	cs.Touch(apue)

	fmt.Println("======= tree Computer_Science=========")
	cs.Tree(1)
	fmt.Println("\n======= tree root=========")
	root.Tree(1)

	fmt.Println("\n======= remove CSAPP.pdf =========")
	cs.Remove(csapp)
	root.Tree(1)
}
