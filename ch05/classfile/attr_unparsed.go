/*
 * @Author: Jiwei
 * @Date: 2022-01-24 17:33:56
 * @LastEditTime: 2022-01-25 09:44:10
 * @LastEditors: Jiwei
 * @Description: \attr_unparsed.go
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_unparsed.go
 *
 */

package classfile

/**
 * @description: 暂时不处理的属性
 * @param {*}
 * @return {*}
 */
type UnparsedAttribute struct {
	name   string
	length uint32
	info   []byte
}

func (self *UnparsedAttribute) readInfo(reader *ClassReader) {
	self.info = reader.readBytes(self.length)
}
