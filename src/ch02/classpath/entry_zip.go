package classpath

import (
	"archive/zip"
	"io/ioutil"
	"path/filepath"
)

type ZipEntry struct {
	//相对路径 压缩文件
	absPath string
}

func newZipEntry(path string) *ZipEntry {
	absPath, err := filepath.Abs(path)
	if err != nil {
		panic(err)
	}
	return &ZipEntry{absPath}
}

func (self *ZipEntry) readClass(className string) ([]byte, Entry, error) {
	r, err := zip.OpenReader(self.absPath)
	if err != nil {
		return nil, nil, err
	}
	defer r.Close()
	for _, f := range r.File {
		if f.Name == className {
			rc, err := f.Open()
			if err != nil {
				return nil, nil, err
			}
			defer rc.Close()
			data, err := ioutil.ReadAll(rc)
			if err != nil {
				return data, nil, nil
			}
			return data, self, nil
		}
	}
	return nil, nil, nil
	//errors.New("class not found :" + className)
}

func (self *ZipEntry) String() string {
	return self.absPath
}
