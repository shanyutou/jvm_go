/*
 * @Author: Jiwei
 * @Date: 2022-01-25 09:19:01
 * @LastEditTime: 2022-01-25 15:11:06
 * @LastEditors: Jiwei
 * @Description:指出源文件名
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_source_file.go
 *
 */
package classfile

/**
 * @description: 源文件名
 * @param {*}
 * @return {*}
 */
type SourceFileAttribute struct {
	cp              ConstantPool
	sourceFileIndex uint16
}

func (self *SourceFileAttribute) readInfo(reader *ClassReader) {
	self.sourceFileIndex = reader.readUint16()
}
func (self *SourceFileAttribute) FileName() string {
	return self.cp.getUtf8(self.sourceFileIndex)
}
