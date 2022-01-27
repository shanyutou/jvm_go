/*
 * @Author: Jiwei
 * @Date: 2022-01-25 15:08:30
 * @LastEditTime: 2022-01-25 15:08:33
 * @LastEditors: Jiwei
 * @Description: attr_exceptions
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_exceptions.go
 *
 */
package classfile

type ExceptionsAttribute struct {
	exceptionIndexTable []uint16
}

func (self *ExceptionsAttribute) readInfo(reader *ClassReader) {
	self.exceptionIndexTable = reader.readUint16s()
}
func (self *ExceptionsAttribute) ExceptionIndexTable() []uint16 {
	return self.exceptionIndexTable
}
