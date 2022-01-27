/*
 * @Author: Jiwei
 * @Date: 2022-01-25 15:11:51
 * @LastEditTime: 2022-01-25 15:11:54
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch03\classfile\attr_markers.go
 *
 */
package classfile

type DeprecatedAttribute struct{ MarkerAttribute }
type SyntheticAttribute struct{ MarkerAttribute }
type MarkerAttribute struct{}

func (self *MarkerAttribute) readInfo(reader *ClassReader) {
	// read nothing
}
