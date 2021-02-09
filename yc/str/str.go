package str

import (
	"github.com/y761350477/go-utils/yc/regex"
	"os"
	"regexp"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/chenhg5/collection"
)

const (
	// 路径分隔符
	PthSep    = string(os.PathSeparator)
	RegexText = `\(.+?\)`
)

// 获取当前条数在第几页（根据每页总条数，当前条数计算）
// num当前条数
// pageTotal每页总条数
func GetPageNum(num, pageTotal int) int {
	i_ := num / pageTotal
	i := num % pageTotal
	if i == 0 {
		return i_
	}
	return i_ + 1
}

// 获取优化后作品的名称
// func GetLabelName(name string) string {
// 	return RegexName(name, `\(.+?\)`)
// }

// 判断slice中是否存在某个item
func IsExistItem(src interface{}, item interface{}) bool {
	return collection.Collect(src).Contains(item)
}

// 复制内容到剪切板
func ClipText(text string) {
	clipboard.WriteAll(text)
}

// 获取文件内容首行信息
func GetContentFirst(text string) string {
	return regex.RegexStr(text, `.*\n`)
}

// 提取番号
func GetSd(src string, RegexText1 string, spli []string) (str_ string) {
	re := regexp.MustCompile(src)
	f := re.FindAllString(RegexText1, -1)
	if len(f) == 1 {
		str_ = f[0]
		return
	} else {
		for _, s_v := range spli {
			for _, f_v := range f {
				if s_v != f_v {
					str_ = f_v
					return
				}
			}
		}
	}
	return
}

// 过滤无效名称
// func RegexName(parentDirName string, RegexLabel string) (str_ string) {
// 	status := regex.RegexStatus(parentDirName, RegexLabel)
// 	if status {
// 		size := regex.RegexStrLen(parentDirName, RegexLabel)
// 		// 保留两个()
// 		i := size - 2
// 		toInt := convert.IntToString(i)
// 		regex := `.*?(\(.*?\)){` + toInt + `}`
// 		str_ = regex.RegexRemove(parentDirName, regex)
// 	}
// 	return
// }

// 根据分隔符分割字符串生成切片
func StrSplit(data, str string) []string {
	result := strings.Split(data, str)
	sle := make([]string, 0, len(result))
	for _, v := range result {
		v = strings.TrimSpace(v)
		sle = append(sle, v)
	}
	return sle
}
