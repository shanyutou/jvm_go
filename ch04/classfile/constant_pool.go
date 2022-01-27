/*
 * @Author: Jiwei
 * @Date: 2022-01-21 15:15:47
 * @LastEditTime: 2022-01-25 15:57:35
 * @LastEditors: Jiwei
 * @Description: 常量池
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\constant_pool.go
 *
 */
package classfile

/**
 * @description: 表头给出的常量池大小比实际大
	1。假设表头给出的值是n，那么常量池的实际大小
	是n–1。第二，有效的常量池索引是1~n–1。0是无
	效索引，表示不指向任何常量。第三，
	CONSTANT_Long_info和CONSTANT_Double_info
	各占两个位置。也就是说，如果常量池中存在这两
	种常量，实际的常量数量比n–1还要少，而且1~n–1
	的某些数也会变成无效索引.
 * @param {*}
 * @return {*}
*/
type ConstantPool []ConstantInfo

/**
 * @description: 读取常量池
 * @param {*ClassReader} reader
 * @return {*}
 */
func readConstantPool(reader *ClassReader) ConstantPool {
	cpCount := int(reader.readUint16())
	cp := make([]ConstantInfo, cpCount)
	for i := 1; i < cpCount; i++ {
		cp[i] = readConstantInfo(reader, cp) // 索引从1开始
		switch cp[i].(type) {
		case *ConstantLongInfo, *ConstantDoubleInfo:
			i++ // 占两个位置
		}

	}
	return cp
}

/**
 * @description: 方法按索引查找常量
 * @param {uint16} index
 * @return {*}
 */
func (self ConstantPool) getConstantInfo(index uint16) ConstantInfo {
	if cpInfo := self[index]; cpInfo != nil {
		return cpInfo
	}
	panic("Invalid constant pool index!")
}

/**
 * @description: 从常量池查找字段或方法的名字和描述符
 * @param {uint16} index
 * @return {*}
 */
func (self ConstantPool) getNameAndType(index uint16) (string, string) {
	ntInfo := self.getConstantInfo(index).(*ConstantNameAndTypeInfo)
	name := self.getUtf8(ntInfo.nameIndex)
	_type := self.getUtf8(ntInfo.descriptorIndex)
	return name, _type
}

/**
 * @description: 从常量池中查找类名
 * @param {uint16} index
 * @return {*}
 */
func (self ConstantPool) getClassName(index uint16) string {
	classInfo := self.getConstantInfo(index).(*ConstantClassInfo)
	return self.getUtf8(classInfo.nameIndex)
}

/**
 * @description: 从常量池查找UTF-8字符串
 * @param {uint16} index
 * @return {*}
 */
func (self ConstantPool) getUtf8(index uint16) string {
	utf8Info := self.getConstantInfo(index).(*ConstantUtf8Info)
	return utf8Info.str
}
