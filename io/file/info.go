package file

import (
	"container/list"
	"github.com/zhang1career/golab/log"
	"io/ioutil"
	"os"
	"path/filepath"
)

func Info(filename string)  {
	fileInfo, err := os.Stat(filename)
	if err != nil {
		log.Error(err.Error())
		return
	}
	
	mt := fileInfo.ModTime()
	
	log.Info("File name:\t\t%s", fileInfo.Name())
	log.Info("Size in bytes:\t%d", fileInfo.Size())
	log.Info("Permissions:\t\t%s", fileInfo.Mode())
	log.Info("Last modified:\t%d-%02d-%02d %02d:%02d:%02d", mt.Year(), mt.Month(), mt.Day(), mt.Hour(), mt.Minute(), mt.Second())
	log.Info("Is Directory:\t%t", fileInfo.IsDir())
}

type FileNode struct {
	FullPath string
	Info os.FileInfo
}

func insertSortedBySize(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		// If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.Size() < element.Value.(FileNode).Info.Size() {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}

func insertSortedByUtime(fileList *list.List, fileNode FileNode) {
	if fileList.Len() == 0 {
		// If list is empty, just insert and return
		fileList.PushFront(fileNode)
		return
	}
	for element := fileList.Front(); element != nil; element = element.Next() {
		if fileNode.Info.ModTime().Before(element.Value.(FileNode).Info.ModTime()) {
			fileList.InsertBefore(fileNode, element)
			return
		}
	}
	fileList.PushBack(fileNode)
}


func InfoListSortedBySize(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("Error reading directory: " + err.Error())
	}
	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			InfoListSortedBySize(fileList, filepath.Join(path, dirFile.Name()))
		} else if dirFile.Mode().IsRegular() {
			insertSortedBySize(fileList, FileNode{FullPath: fullpath, Info: dirFile})
		}
	}
}

func InfoListSortedByUtime(fileList *list.List, path string) {
	dirFiles, err := ioutil.ReadDir(path)
	if err != nil {
		log.Error("Error reading directory: " + err.Error())
	}
	for _, dirFile := range dirFiles {
		fullpath := filepath.Join(path, dirFile.Name())
		if dirFile.IsDir() {
			InfoListSortedByUtime(fileList, filepath.Join(path, dirFile.Name()))
		} else if dirFile.Mode().IsRegular() {
			insertSortedByUtime(fileList, FileNode{FullPath: fullpath, Info: dirFile})
		}
	}
}
