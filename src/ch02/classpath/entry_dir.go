package classpath

import (
	"io/ioutil"
	"path/filepath"
)

type DirEntry struct {
	//文件目录 go 结构体不需要显示实现接口，只要方法匹配即可。go没有专门的构造函数
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
