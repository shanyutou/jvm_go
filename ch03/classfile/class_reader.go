/*
 * @Author: Jiwei
 * @Date: 2022-01-20 17:00:59
 * @LastEditTime: 2022-01-21 16:49:04
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \jvmgo\ch03\classfile\class_reader.go
 *
 */
package classfile

import (
	"encoding/binary"
)

type ClassReader struct {
	data []byte
}

func (self *ClassReader) readUint8() uint8 {
	val := self.data[0]
	self.data = self.data[1:] // reslice语法
	return val
}

func (self *ClassReader) readUint16() uint16 {
	val := binary.BigEndian.Uint16(self.data)
	self.data = self.data[2:] // reslice语法
	return val
} // u2
func (self *ClassReader) readUint32() uint32 {
	val := binary.BigEndian.Uint32(self.data)
	self.data = self.data[4:] // reslice语法
	return val
} // u4
func (self *ClassReader) readUint64() uint64 {
	val := binary.BigEndian.Uint64(self.data)
	self.data = self.data[8:] // reslice语法
	return val
}
func (self *ClassReader) readUint16s() []uint16 {
	n := self.readUint16() //读取uint16表，表的大小由开头的uint16数据指出
	s := make([]uint16, n)
	for i := range s {
		s[i] = self.readUint16()
	}
	return s
}
func (self *ClassReader) readBytes(length uint32) []byte {
	bytes := self.data[:length]
	self.data = self.data[length:]
	return bytes
}
