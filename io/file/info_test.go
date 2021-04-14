package file_test

import (
	"container/list"
	"fmt"
	"github.com/zhang1career/golab/io/file"
	"testing"
)

func TestInfo(t *testing.T) {
	file.Info("test.txt")
}

func TestInfoListSortedBySize(t *testing.T) {
	fileList := list.New()
	file.InfoListSortedBySize(fileList, "/Users/zhang/Downloads")
	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Printf("%d %s\n", element.Value.(file.FileNode).Info.Size(), element.Value.(file.FileNode).FullPath)
	}
}

func TestInfoListSortedByUtime(t *testing.T) {
	fileList := list.New()
	file.InfoListSortedByUtime(fileList, "/Users/zhang/Downloads")
	for element := fileList.Front(); element != nil; element = element.Next() {
		fmt.Print(element.Value.(file.FileNode).Info.ModTime())
		fmt.Printf("%s\n", element.Value.(file.FileNode).FullPath)
	}
}
