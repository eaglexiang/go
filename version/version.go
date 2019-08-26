/*
 * @Author: EagleXiang
 * @Github: https://github.com/eaglexiang
 * @Date: 2019-01-02 15:15:55
 * @LastEditors: EagleXiang
 * @LastEditTime: 2019-02-19 15:37:01
 */

package version

import (
	"errors"
	"strconv"
	"strings"
)

// 比较结果的三种基本状态
const (
	Greater = iota
	Less
	Equal
)

// Version 版本，包含版本号字符串与其对应的int数组
type Version struct {
	nodes []uint
	Raw   string
}

// CreateVersion 根据格式由字符串创建Version
func CreateVersion(src string) (Version, error) {
	result := Version{Raw: src}
	items := strings.Split(result.Raw, ".")
	result.nodes = make([]uint, len(items))
	for index, item := range items {
		itemInt, err := strconv.ParseUint(item, 10, 32)
		if err != nil {
			return result, errors.New("invalid version string")
		}
		result.nodes[index] = uint(itemInt)
	}
	return result, nil
}

// IsGreaterThan src >= des
func (src Version) IsGreaterThan(des Version) bool {
	return src.compareWith(des) == Greater
}

// IsLessThan src <= des
func (src Version) IsLessThan(des Version) bool {
	return src.compareWith(des) == Less
}

// Equals src与des相等
func (src Version) Equals(des Version) bool {
	return src.compareWith(des) == Equal
}

// IsGTOrE2 sr >= des
func (src Version) IsGTOrE2(des Version) bool {
	relation := src.compareWith(des)
	return relation == Greater || relation == Equal
}

// IsLTOrE2 sr <= des
func (src Version) IsLTOrE2(des Version) bool {
	relation := src.compareWith(des)
	return relation == Less || relation == Equal
}

func (src Version) compareWith(des Version) int {
	maxLen := len(src.nodes)
	if desLen := len(des.nodes); desLen > maxLen {
		maxLen = desLen
	}
	// 补齐较短的版本号
	for len(src.nodes) < maxLen {
		src.nodes = append(src.nodes, 0)
	}
	for len(des.nodes) < maxLen {
		des.nodes = append(des.nodes, 0)
	}

	// 从前往后进行比较
	for i, v := range src.nodes {
		if v < des.nodes[i] {
			return Less
		} else if v > des.nodes[i] {
			return Greater
		}
	}
	return Equal
}
