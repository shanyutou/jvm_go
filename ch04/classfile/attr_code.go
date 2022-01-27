/*
 * @Author: Jiwei
 * @Date: 2022-01-25 10:06:38
 * @LastEditTime: 2022-01-25 15:37:18
 * @LastEditors: Jiwei
 * @Description: attr_code.go 字节码等
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_code.go
 *
 */
package classfile

type CodeAttribute struct {
	cp             ConstantPool
	maxStack       uint16                 // 栈的最大深度
	maxLocals      uint16                 // 局部变量表大小
	code           []byte                 // 字节码
	exceptionTable []*ExceptionTableEntry // 异常表
	attributes     []AttributeInfo        // 属性表
}

type ExceptionTableEntry struct {
	startPc   uint16
	endPc     uint16
	handlerPc uint16
	catchType uint16
}

func (self *CodeAttribute) readInfo(reader *ClassReader) {
	self.maxStack = reader.readUint16()
	self.maxLocals = reader.readUint16()
	codeLength := reader.readUint32()
	self.code = reader.readBytes(codeLength)
	self.exceptionTable = readExceptionTable(reader)
	self.attributes = readAttributes(reader, self.cp)
}

func readExceptionTable(reader *ClassReader) []*ExceptionTableEntry {
	exceptionTableLength := reader.readUint16()
	exceptionTable := make([]*ExceptionTableEntry, exceptionTableLength)
	for i := range exceptionTable {
		exceptionTable[i] = &ExceptionTableEntry{
			startPc:   reader.readUint16(),
			endPc:     reader.readUint16(),
			handlerPc: reader.readUint16(),
			catchType: reader.readUint16(),
		}

	}
	return exceptionTable
}
