/*
 * @Author: Jiwei
 * @Date: 2022-01-25 11:34:24
 * @LastEditTime: 2022-01-25 15:06:57
 * @LastEditors: Jiwei
 * @Description: attr_line_number_table.go e属性表存放方法的行号信息
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_line_number_table.go
 *
 */
package classfile

type LineNumberTableAttribute struct {
	lineNumberTable []*LineNumberTableEntry
}
type LineNumberTableEntry struct {
	startPc    uint16
	lineNumber uint16
}

func (self *LineNumberTableAttribute) readInfo(reader *ClassReader) {
	lineNumberTableLength := reader.readUint16()
	self.lineNumberTable = make([]*LineNumberTableEntry, lineNumberTableLength)
	for i := range self.lineNumberTable {
		self.lineNumberTable[i] = &LineNumberTableEntry{
			startPc:    reader.readUint16(),
			lineNumber: reader.readUint16(),
		}
	}
}
