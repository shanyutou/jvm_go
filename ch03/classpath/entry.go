/*
 * @Author: Jiwei
 * @Date: 2022-01-14 15:20:14
 * @LastEditTime: 2022-01-18 10:18:34
 * @LastEditors: Jiwei
 * @Description:entry
 * @FilePath: \Go_WorkSpace\src\jvmgo\ch02\classpath\entry.go
 *
 */
package classpath

import (
	"os"
	"strings"
)

const pathListSeparator = string(os.PathListSeparator)

type Entry interface {
	readClass(className string) ([]byte, Entry, error)
	String() string
}

func newEntry(path string) Entry {
	if strings.Contains(path, pathListSeparator) {
		return nil
	}

	return nil

}
