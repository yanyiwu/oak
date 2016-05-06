package util

import (
	"io/ioutil"
	"os"

	"github.com/golang/glog"
)

func GetDirsFromDir(dirpath string) []string {
	results := make([]string, 0)
	files, err := ioutil.ReadDir(dirpath)
	if err != nil {
		panic(err)
	}
	for _, f := range files {
		if f.IsDir() {
			results = append(results, f.Name())
		}
	}
	return results
}

func CreateDirIfNotExists(dirpath string) {
	_, err := os.Stat(dirpath)
	if !os.IsExist(err) {
		glog.Info("create dir:", dirpath)
		os.MkdirAll(dirpath, os.ModePerm)
	}
}
