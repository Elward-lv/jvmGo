package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	/**
	 *	路劲之间使用/分割 文件名有class后缀s 如Java/lang/Object.class
	 *  返回是读取到的字节数据(go 允许返回多个值)
	 **/
	readClass(className string) ([]byte, Entry, error)
	String() string
}

// 根据入参创建不同类型的entry
func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return newCompositeEntry(path)
	}
	if strings.HasSuffix(path, "*") {
		return newWildcardEntry(path)
	}
	if strings.HasSuffix(path, ".jar") || strings.HasSuffix(path, ".JAR") || strings.HasSuffix(path, ".zip") || strings.HasSuffix(path, ".ZIP") {
		return newZipEntry(path)
	}
	return newDirEntry(path)
}
