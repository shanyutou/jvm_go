/*
 * @Author: Jiwei
 * @Date: 2022-01-26 10:09:31
 * @LastEditTime: 2022-01-26 15:53:52
 * @LastEditors: Jiwei
 * @Description:Thread.go
 * @FilePath: \jvmgo\ch04\rtda\thread.go
 *
 */
package rtda

type Thread struct {
	pc    int
	stack *Stack
}

func NewThread() *Thread {
	return &Thread{
		stack: newStack(1024),
	}
}
func (self *Thread) PC() int      { return self.pc } // getter
func (self *Thread) SetPC(pc int) { self.pc = pc }   // setter
func (self *Thread) PushFrame(frame *Frame) {
	self.stack.push(frame)
}
func (self *Thread) PopFrame() *Frame {
	return self.stack.pop()
}
func (self *Thread) CurrentFrame() *Frame {
	return self.stack.top()
}
