/*
 * @Author: Jiwei
 * @Date: 2022-01-24 16:13:45
 * @LastEditTime: 2022-01-24 16:16:19
 * @LastEditors: Jiwei
 * @Description: cp_class.go
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\cp_class.go
 *
 */

package classfile

type ConstantClassInfo struct {
	cp        ConstantPool
	nameIndex uint16
}

func (self *ConstantClassInfo) readInfo(reader *ClassReader) {
	self.nameIndex = reader.readUint16()
}
func (self *ConstantClassInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
