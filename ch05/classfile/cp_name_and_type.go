/*
 * @Author: Jiwei
 * @Date: 2022-01-24 10:02:30
 * @LastEditTime: 2022-01-24 10:02:34
 * @LastEditors: Jiwei
 * @Description:cp_name_and_type.go
 * @FilePath: \jvmgo\ch03\classfile\cp_name_and_type.go
 *
 */
package classfile

type ConstantNameAndTypeInfo struct {
	nameIndex       uint16
	descriptorIndex uint16
}

func (self *ConstantNameAndTypeInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
	self.descriptorIndex = reader.readUint16()
}
