package fs

import (
	"io/ioutil"
	"os"
)

type filesInfoInDirectory = func(dirname string) ([]os.FileInfo, error)

type fileReader interface {
	ReadFile(filename string) ([]byte, error)
	OpenFile(name string, flag int, perm os.FileMode) (*os.File, error)
}

type dirReader interface {
	ReadDir(dirPath string) ([]os.FileInfo, error)
	CreateDir(dirPath string) bool
	ExistsDir(dirPath string) bool
}

type fsReader struct {
	fileReader
	dirReader
}

func (fsReader) OpenFile(name string, flag int, perm os.FileMode) (*os.File, error) {
	return os.OpenFile(name, flag, perm)
}

func (fsReader) ReadFile(filename string) ([]byte, error) {
	return ioutil.ReadFile(filename)
}

func (fsReader) ReadDir(dirPath string) ([]os.FileInfo, error) {
	return ioutil.ReadDir(dirPath)
}

func (fsReader) CreateDir(dirPath string) bool {
	return os.MkdirAll(dirPath, os.ModePerm) == nil
}

func (fsReader) ExistsDir(dirPath string) bool {
	_, err := os.Stat(dirPath)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return true
}

func NewFsReader() *fsReader {
	return &fsReader{}
}
