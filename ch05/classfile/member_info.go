/*
 * @Author: Jiwei
 * @Date: 2022-01-21 14:54:36
 * @LastEditTime: 2022-01-24 15:28:50
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\member_info.go
 *
 */
package classfile

type MemberInfo struct {
	cp              ConstantPool
	accessFlags     uint16
	nameIndex       uint16
	descriptorIndex uint16
	attributes      []AttributeInfo
}

func readMembers(reader *ClassReader, cp ConstantPool) []*MemberInfo {
	numberCount := reader.readUint16()
	members := make([]*MemberInfo, numberCount)
	for i := range members {
		members[i] = readMember(reader, cp)
	}
	return members

}
func readMember(reader *ClassReader, cp ConstantPool) *MemberInfo {
	return &MemberInfo{
		cp:              cp,
		accessFlags:     reader.readUint16(),
		nameIndex:       reader.readUint16(),
		descriptorIndex: reader.readUint16(),
		attributes:      readAttributes(reader, cp),
	}
}
func (self *MemberInfo) AccessFlags() uint16 {
	return self.accessFlags
} // getter
func (self *MemberInfo) Name() string {
	return self.cp.getUtf8(self.nameIndex)
}
func (self *MemberInfo) Descriptor() string {
	return self.cp.getUtf8(self.descriptorIndex)
}
