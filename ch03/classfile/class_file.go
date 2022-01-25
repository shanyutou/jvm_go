/*
 * @Author: Jiwei
 * @Date: 2022-01-20 17:40:44
 * @LastEditTime: 2022-01-25 15:32:14
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\class_file.go
 *
 */
package classfile

import "fmt"

type ClassFile struct {
	magic        uint32          // 魔数
	minorVersion uint16          // 小版本号
	majorVersion uint16          // 大版本号
	constantPool ConstantPool    // 常量池
	accessFlags  uint16          // 访问标志
	thisClass    uint16          // 类名称
	superClass   uint16          // 超类名称
	interfaces   []uint16        // 接口名称
	fields       []*MemberInfo   // 字段表
	methods      []*MemberInfo   // 方法表
	attributes   []AttributeInfo // 属性表
}

/**
 * @description: 数把[]byte解析成ClassFile结构体
 * @param {[]byte} classData
 * @return {*}
 */
func Parse(classData []byte) (cf *ClassFile, err error) {
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

/**
 * @description: 依次调用其他方法解析class文件
 * @param {*ClassReader} reader
 * @return {*}
 */
func (self *ClassFile) read(reader *ClassReader) {
	self.readAndCheckMagic(reader)
	self.readAndCheckVersion(reader)
	self.constantPool = readConstantPool(reader)
	self.accessFlags = reader.readUint16()
	self.thisClass = reader.readUint16()
	self.superClass = reader.readUint16()
	self.interfaces = reader.readUint16s()
	self.fields = readMembers(reader, self.constantPool)
	self.methods = readMembers(reader, self.constantPool)
	self.attributes = readAttributes(reader, self.constantPool)
}

/**
 * @description: 暴露结构体的字段
 * @param {*}
 * @return {*}
 */
func (self *ClassFile) MajorVersion() uint16 {
	return self.majorVersion
}

func (self *ClassFile) MinorVersion() uint16 {
	return self.minorVersion
}

/**
 * @description: 从常量池中找超类名
 * @param {*}
 * @return {*}
 */
func (self *ClassFile) SuperClassName() string {
	return self.constantPool.getClassName(self.superClass)
}

/**
 * @description: 从常量池查找接口名
 * @param {*}
 * @return {*}
 */
func (self *ClassFile) InterfaceNames() []string {
	interfaceNames := make([]string, len(self.interfaces))
	for i, cpIndex := range self.interfaces {
		interfaceNames[i] = self.constantPool.getClassName(cpIndex)
	}
	return interfaceNames
}

/**
 * @description: 检查
 * @param {*ClassReader} reader
 * @return {*}
 */
func (self *ClassFile) readAndCheckMagic(reader *ClassReader) {
	magic := reader.readUint32()
	if magic != 0xCAFEBABE {
		panic("java.lang.ClassFormatError: magic!")
	}
}

/**
 * @description: 检查版本号
 * @param {*ClassReader} read
 * @return {*}
 */
func (self *ClassFile) readAndCheckVersion(read *ClassReader) {
	self.minorVersion = read.readUint16()
	self.majorVersion = read.readUint16()
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

func (self *ClassFile) ConstantPool() ConstantPool {
	return self.constantPool
} // getter
func (self *ClassFile) AccessFlags() uint16 {
	return self.accessFlags
} // getter
func (self *ClassFile) Fields() []*MemberInfo {
	return self.fields
} // getter
func (self *ClassFile) Methods() []*MemberInfo {
	return self.methods
} // getter
func (self *ClassFile) ClassName() string {
	return self.constantPool.getClassName(self.thisClass)
}
