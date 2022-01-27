/*
 * @Author: Jiwei
 * @Date: 2022-01-26 10:33:38
 * @LastEditTime: 2022-01-26 15:58:33
 * @LastEditors: Jiwei
 * @Description: jvm_stack.go
 * @FilePath: \jvmgo\ch04\rtda\jvm_stack.go
 *
 */

package rtda

type Stack struct {
	maxSize uint
	size    uint
	_top    *Frame
}

func newStack(maxSize uint) *Stack {
	return &Stack{
		maxSize: maxSize,
	}
}
func (self *Stack) push(frame *Frame) {
	if self.size > self.maxSize {
		panic("java.lang.StackOverflowError")
	}
	if self._top != nil {
		frame.lower = self._top
	}
	self._top = frame
	self.size++
}
func (self *Stack) pop() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	top := self._top
	self._top = top.lower
	top.lower = nil
	self.size--
	return top
}
func (self *Stack) top() *Frame {
	if self._top == nil {
		panic("jvm stack is empty!")
	}
	return self._top
}
