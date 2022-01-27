/*
 * @Author: Jiwei
 * @Date: 2022-01-24 16:01:48
 * @LastEditTime: 2022-01-25 09:44:05
 * @LastEditors: Jiwei
 * @Description: cp_string.go
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\cp_string.go
 *
 */
package classfile

type ConstantStringInfo struct {
	cp          ConstantPool
	stringIndex uint16
}

func (self *ConstantStringInfo) readInfo(reader *ClassReader) {
	self.stringIndex = reader.readUint16()
}
func (self *ConstantStringInfo) String() string {
	return self.cp.getUtf8(self.stringIndex)
}
