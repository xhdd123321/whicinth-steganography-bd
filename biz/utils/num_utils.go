package utils

import "strconv"

// IsNum 判断字符串是否是数字
func IsNum(s string) bool {
	_, err := strconv.Atoi(s)
	return err == nil
}
