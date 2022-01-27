/*
 * @Author: Jiwei
 * @Date: 2022-01-18 09:48:28
 * @LastEditTime: 2022-01-21 11:44:36
 * @LastEditors: Jiwei
 * @Description:
 * @FilePath: \jvmgo\ch03\classpath\entry_composite.go
 *
 */
package classpath

import (
	"errors"
	"strings"
)

type CompositeEntry []Entry

func newCompositeEntry(pathList string) CompositeEntry {
	return nil
}
func (self CompositeEntry) readClass(className string) ([]byte, Entry, error) {
	for _, entry := range self {
		data, from, err := entry.readClass(className)
		if err == nil {
			return data, from, nil
		}
	}
	return nil, nil, errors.New("class not found: " + className)
}
func (self CompositeEntry) String() string {
	strs := make([]string, len(self))
	for i, entry := range self {
		strs[i] = entry.String()
	}
	return strings.Join(strs, pathListSeparator)
}
