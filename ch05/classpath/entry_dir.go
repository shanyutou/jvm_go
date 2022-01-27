/*
 * @Author: Jiwei
 * @Date: 2022-01-14 15:58:17
 * @LastEditTime: 2022-01-17 17:30:34
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch02\classpath\entry_dir.go
 *
 */
package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	absDir string
}

func newDirEntry(path string) *DirEntry {
	absDir, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &DirEntry{absDir}
}
func (self *DirEntry) readClass(className string) ([]byte, Entry, error) {
	fileName := filepath.Join(self.absDir, className)
	data, err := ioutil.ReadFile(fileName)
	return data, self, err
}

func (self *DirEntry) String() string {
	return self.absDir
}
