// BUG(polaris): 一个段落英文开头的大小写转换有问题，比如 go中文网 中的 go 不会转为 Go。

package autocorrect

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

var otherDicts = make(map[string]string)

// AddDict 支持自定义添加字典
func AddDict(dict map[string]string) {
	for k, v := range dict {
		otherDicts[k] = v
	}
}

// AutoSpace 自动给中英文之间加上空格
func AutoSpace(str string) string {
	out := ""
	for _, r := range str {
		out = addSpaceAtBoundary(out, r)
	}

	return out
}

// AutoCorrect 对常见英文单词进行大家一般写法的纠正，如 go -> Go
func AutoCorrect(str string) string {
	// oldNews := make([]string, 2*(len(dicts)+len(otherDicts)))
	// oldNews := make([]string, 2*(len(dicts)+len(otherDicts)))
	var oldNews []string
	for from, to := range dicts {
		if !strings.HasPrefix(from, "`") {
			oldNews = append(oldNews, " "+from+" ")
		}

		if !strings.HasPrefix(to, "`") {
			oldNews = append(oldNews, " "+to+" ")
		}
	}

	replacer := strings.NewReplacer(oldNews...)
	return replacer.Replace(str)
}

// Convert 先执行 AutoSpace，然后执行 AutoCorrect
// 由于处理专有名词可能会出现替换命令中的单词，默认不开启
func Convert(str string, b ...bool) string {
	var bo bool
	switch len(b) {
	case 0:
		bo = false
	case 1:
		bo = b[0]
	default:
		panic("too many parameters")
	}

	if bo {
		return AutoCorrect(AutoSpace(str))
	}
	return AutoSpace(str)
}

func addSpaceAtBoundary(prefix string, nextChar rune) string {
	if len(prefix) == 0 {
		return string(nextChar)
	}

	r, size := utf8.DecodeLastRuneInString(prefix)
	if isLatin(size) != isLatin(utf8.RuneLen(nextChar)) &&
		isAllowSpace(nextChar) && isAllowSpace(r) {
		return prefix + " " + string(nextChar)
	}

	return prefix + string(nextChar)
}

func isLatin(size int) bool {
	return size == 1
}

func isAllowSpace(r rune) bool {
	return !unicode.IsSpace(r) && !unicode.IsPunct(r)
}
