/*
 * @Author: Jiwei
 * @Date: 2022-01-24 16:35:34
 * @LastEditTime: 2022-01-24 17:06:40
 * @LastEditors: Jiwei
 * @Description: cp_member_ref.go
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\cp_member_ref.go
 *
 */
package classfile

/**
 * @description: 字段符号引用
 * @param {*}
 * @return {*}
 */
type ConstantFieldrefInfo struct{ ConstantMemberrefInfo }

/**
* @description: 普通（非接口）方法符号引用
* @param {*}
* @return {*}
 */
type ConstantMethodrefInfo struct{ ConstantMemberrefInfo }

/**
* @description: 示接口方法符号引用
* @param {*}
* @return {*}
 */
type ConstantInterfaceMethodrefInfo struct{ ConstantMemberrefInfo }

/**
 * @description: 统一的结构体ConstantMemberrefInfo
 * @param {*}
 * @return {*}
 */
type ConstantMemberrefInfo struct {
	cp               ConstantPool
	classIndex       uint16
	nameAndTypeIndex uint16
}

func (self *ConstantMemberrefInfo) readInfo(reader *ClassReader) {
	self.classIndex = reader.readUint16()
	self.nameAndTypeIndex = reader.readUint16()
}
func (self *ConstantMemberrefInfo) ClassName() string {
	return self.cp.getUtf8(self.classIndex)
}

func (self *ConstantMemberrefInfo) nameAndDescriptor() (string, string) {
	return self.cp.getNameAndType(self.nameAndTypeIndex)

}
