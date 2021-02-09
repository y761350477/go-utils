package regex

import (
	"regexp"
)

// 正则匹配
func RegexStr(text, regex string) string {
	re := regexp.MustCompile(regex)
	return re.FindString(text)
}

// 获取子匹配中的首个匹配
func RegexStrFirst(text, regex string) string {
	re := regexp.MustCompile(regex)
	match := re.FindAllStringSubmatch(text, -1)
	return match[0][0]
}

// 获取匹配到的个数
func RegexStrLen(text, regex string) int {
	re := regexp.MustCompile(regex)
	match := re.FindAllStringSubmatch(text, -1)
	return len(match)
}

// 删除正则匹配的内容
func RegexRemove(text, regex string) string {
	re := regexp.MustCompile(regex)
	return re.ReplaceAllString(text, ``)
}

func RegexStatus(text, regex string) bool {
	re := regexp.MustCompile(regex)
	findString := re.FindString(text)
	if findString != `` {
		return true
	}
	return false
}
