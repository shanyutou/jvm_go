/*
 * @Author: Jiwei
 * @Date: 2022-01-26 11:14:06
 * @LastEditTime: 2022-01-26 16:27:15
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \jvmgo\ch04\rtda\slot.go
 *
 */

/**
 * @description: 局部变量表中存放内容
 * @param {*}
 * @return {*}
 */
package rtda

type Slot struct {
	num int32
	ref *Object
}
