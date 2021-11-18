package Composite_Pattern

import (
	"container/list"
	"fmt"
	"strings"
)

/**
 * 实现一个文件管理工具，支持对文件夹和文件的创建和删除功能，并且提供 tree 接口显示文件系统结构。
 * 设计类图如下：https://gitee.com/jakel-in/images/raw/master/2021-11/Composite-Pattern-Example.png
 */

// 抽象部件：抽象文件结点
type IFile interface {
	Touch(file IFile)
	Remove(file IFile)
	Tree(depth int)
}

// 容器部件：文件夹
type Folder struct {
	fileName string
	childs   list.List
}

func (f *Folder) Touch(file IFile) {
	f.childs.PushBack(file)
}
func (f *Folder) Remove(file IFile) {
	for element := f.childs.Front(); element != nil; element = element.Next() {
		if element.Value == file {
			f.childs.Remove(element)
		}
	}
}
func (f *Folder) Tree(depth int) {
	fmt.Println(strings.Repeat("-", depth), f.fileName)
	for element := f.childs.Front(); element != nil; element = element.Next() {
		file, _ := element.Value.(IFile)
		if file != nil {
			file.Tree(depth + 1)
		}
	}
}

// 叶子：文件
type File struct {
	fileName string
}

func (f *File) Touch(file IFile) {
	fmt.Println("Error! Cannot touch from a file.")
}
func (f *File) Remove(file IFile) {
	fmt.Println("Error! Cannot remove from a file.")
}
func (f *File) Tree(depth int) {
	fmt.Println(strings.Repeat("-", depth), f.fileName)
}
