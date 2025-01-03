package classfile

// 属性表
/**
23种预定义可以分为三组：
	1. 实现Java虚拟机所必须的 共5种
	2. Java类库所必须的 共有12种
	3. 提供给工具使用的 共有6种。可选的 可以不出现在Class文件中 比如LineNumberTable属性 在异常堆栈之中显示行号

jdk 1.0 只有六种定义
jdk 1.1 增加了3种
j2se 5.0 增加了9种，主要用于支持反省和注解
Java se 6增加了StackMapTable属性 用于优化字节码验证
Java se 7 增加了BootstrapMethods属性，用于支持新增的invoke dynamic指令。
Java se8 又增加了3种属性。具体可以参考P117（自己动手写Java虚拟机）
*/
type AttributeInfo interface {
	readInfo(reader *ClassReader)
}

func newAttributeInfo(attrName string, attrLen uint32, cp ConstantPool) AttributeInfo {
	switch attrName {
	case "Code":
		return &CodeAttribute{cp: cp}
	case "ConstantValue":
		return &ConstantValueAttribute{}
	case "Deprecated":
		return &DeprecatedAttribute{}
	case "Exceptions":
		return &ExceptionAttribute{}
	case "LineNumberTable":
		return &LineNumberTableAttribute{}
	case "LocalVariableTable":
		return &LocalVariableTableAttribute{}
	case "SourceFile":
		return &SourceFileAttribute{}
	case "Synthetic":
		return &SyntheticAttribute{}
	default:
		return &UnparsedAttribute{attrName, attrLen, nil}
	}
}
