package classfile

//常量池的常量分为:
//1. 字面量： 数字常量 和 字符串常量
//2. 符号引用: 类、接口名、字段和方法信息等。
//	除了字面量 其他常量都是通过索引直接或者间接指向ConstantUtf8Info常量
/**
 *		比如Constant_Fieldref_info:
 * 		Fieldref_info   -|class_index					class_info
 *														-|name_index			Utf8_info  -|bytes
 *														-|NameAndType_info
 *						-|name_and_type_index			NameAndType_info
 *														-|name_index				Utf8_info  -|bytes
 *														-|descriptor_index
**/

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	//自定义MUTF-8编码
	self.str = decodeMUTF8(bytes)
}

func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
