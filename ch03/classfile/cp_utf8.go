/*
 * @Author: Jiwei
 * @Date: 2022-01-21 17:20:57
 * @LastEditTime: 2022-01-21 17:23:16
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \jvmgo\ch03\classfile\cp_utf8.go
 *
 */
package classfile

type ConstantUtf8Info struct {
	str string
}

func (self *ConstantUtf8Info) readInfo(reader *ClassReader) {
	length := uint32(reader.readUint16())
	bytes := reader.readBytes(length)
	self.str = decodeMUTF8(bytes)
}

// todo 简化版，完整版请阅读本章源代码
func decodeMUTF8(bytes []byte) string {
	return string(bytes)
}
