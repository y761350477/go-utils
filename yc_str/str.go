package yc_str

import (
	"github.com/y761350477/go-utils/yc_regex"
	"os"
	"strings"

	"github.com/atotto/clipboard"
	"github.com/chenhg5/collection"
)

const (
	// 路径分隔符
	PthSep = string(os.PathSeparator)
	Debug  = "Debug"
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

// 判断slice中是否存在某个item
func IsExistItem(src interface{}, item interface{}) bool {
	return collection.Collect(src).Contains(item)
}

// 复制内容到剪切板
func ClipText(text string) {
	clipboard.WriteAll(text)
}

// 获取文件内容首行信息
func GetContentFirstLine(text string) string {
	return yc_regex.Match(text, yc_regex.FirstLineFilter)
}

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
