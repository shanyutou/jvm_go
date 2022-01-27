/*
 * @Author: Jiwei
 * @Date: 2022-01-25 09:33:39
 * @LastEditTime: 2022-01-25 09:44:40
 * @LastEditors: Jiwei
 * @Description: 表示常量
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_constant_value.go
 *
 */
/**
 * @description:
 * @param {*}
 * @return {*}
 */
package classfile

type ConstantValueAttribute struct {
	constantValueIndex uint16
}

func (self *ConstantValueAttribute) readInfo(reader *ClassReader) {
	self.constantValueIndex = reader.readUint16()
}
func (self *ConstantValueAttribute) ConstantValueIndex() uint16 {
	return self.constantValueIndex
}
