package classfile

import "fmt"

type ClassFile struct {
	//magic uint32
	minorVersion uint16
	majorVersion uint16
	//常量池
	constantPool ConstantPool
	accessFlags  uint16
	thisClass    uint16
	superClass   uint16
	interfaces   []uint16
	fields       []*MemberInfo
	methods      []*MemberInfo
	attributes   []AttributeInfo
}

func Parse(classData []byte) (cf *ClassFile, err error) {
	//go没有异常处理机制 只有panic 和 recover机制
	defer func() {
		if r := recover(); r != nil {
			var ok bool
			err, ok = r.(error)
			if !ok {
				err = fmt.Errorf("%v", r)
			}
		}
	}()
	cr := &ClassReader{classData}
	cf = &ClassFile{}
	cf.read(cr)
	return
}

func (self *ClassFile) read(reader *ClassReader) {
	//1.检查magic
	self.readAndCheckMagic(reader)
	//2.检查version
	self.readAndCheckVersion(reader)
	//3.常量池
	self.constantPool = readConstantPool(reader)
	//4.类访问标识 16位 标识class是类还是接口 访问级别是public还是private等
	self.accessFlags = reader.readUint16()
	//5.类 完全限定名 b/a/c
	self.thisClass = reader.readUint16()
	//6.超类
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

func readAttributes(reader *ClassReader, pool ConstantPool) []AttributeInfo {

}

func readMembers(reader *ClassReader, pool ConstantPool) []*MemberInfo {

}

func readConstantPool(reader *ClassReader) ConstantPool {

}

// 1.检查class文件的magic是否合法
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: Wrong magic!")
	}
}

// 小写即为内部访问 包内
// 2. 检查class的版本号
func (self *ClassFile) readAndCheckVersion(reader *ClassReader) {
	self.minorVersion = reader.readUint16()
	self.majorVersion = reader.readUint16()
	switch self.majorVersion {
	case 45:
		return
	case 46, 47, 48, 49, 50, 51, 52:
		if self.minorVersion == 0 {
			return
		}
	}
	panic("java.lang.UnsupportedClassVersionError!")
}

// 大写即为公开访问
func (self *ClassFile) MinorVersion() uint16 { //getter
	return self.minorVersion
}

func (self *ClassFile) MajorVersion() uint16 { //getter
	return self.majorVersion
}

func (self *ClassFile) ConstantPool() ConstantPool { //getter

}

func (self *ClassFile) AccessFlags() uint16 { //getter

}

func (self *ClassFile) Fields() []*MemberInfo { //getter

}

func (self *ClassFile) Methods() []*MemberInfo { //getter

}

func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}

func (self *ClassFile) SuperClassName() string {
	if self.superClass > 0 {
		return self.constantPool.getClassName(self.superClass)
	}
	return "" //只有Object超类
}

func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		//从常量池之中查找接口名称
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}
