/*
 * @Author: Jiwei
 * @Date: 2022-01-26 10:34:14
 * @LastEditTime: 2022-01-26 17:18:36
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \jvmgo\ch04\rtda\frame.go
 *
 */
package rtda

type Frame struct {
	lower        *Frame        //实现链表数据结构
	localVars    LocalVars     //保存局部变量表
	operandStack *OperandStack //保存局部变量表
}

func NewFrame(maxLocals, maxStack uint) *Frame {
	return &Frame{
		localVars:    newLocalVars(maxLocals),
		operandStack: newOperandStack(maxStack),
	}
}

func (self *Frame) LocalVars() LocalVars {
	return self.localVars
}

func (self *Frame) OperandStack() *OperandStack {
	return self.operandStack
}
