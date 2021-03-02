package yc_regex

import (
	"github.com/dlclark/regexp2"
	"regexp"
	"strings"
)

const (
	// 匹配作品名称如：「名稱： (HD1080P H264)(DANDY)(1dandy...)宿泊ドックの数日...」
	C080TitleFilter = `.*名稱：.*`

	// 提取 「名稱： (HD1080P H264)(DANDY)(1dandy...)宿泊ドックの数日...」 的信息 「(DANDY)(1dandy...)宿泊ドックの数日...」
	C080TitleMainInfoFilter = `(\([^\(\)]+?\)){2}[^\(]+$`
)

const (
	FirstLineFilter = `.*\n`
)

const (
	TextTypeFilter = `.txt`
	RarTypeFilter  = `.rar`
	SrtTypeFilter  = `.srt`
)

const (
	Ed2kFilter    = `^ed2k://.*`    // ed2k:// 打头的磁力
	RarFilter     = `.*\.rar$`      // .rar 后缀的压缩文件
	Mp4Filter     = `.*\.mp4$`      // .mp4 后缀的视频文件
	WmvFilter     = `.*\.wmv$`      // .wmv 后缀的视频文件
	ThunderFilter = `^thunder://.*` // thunder:// 打头的磁力
	MagnetFilter  = `^magnet:?.*`   // magnet:? 打头的磁力
)

// 正则匹配
func Match(s, regex string) string {
	var re *regexp.Regexp

	re = regexp.MustCompile(regex)
	return re.FindString(s)
}

func MatchExist(s, regex string) bool {
	return Match(s, regex) != ""
}

func MatchSub(s, regex string) [][]string {
	var re *regexp.Regexp

	re = regexp.MustCompile(regex)
	// 只匹配第一个满足的条件
	return re.FindAllStringSubmatch(s, -1)
}

// 删除正则匹配的内容
func RemoveFirstMatch(src, regex string) string {
	var (
		re         *regexp.Regexp
		findString string
	)

	re = regexp.MustCompile(regex)
	findString = re.FindString(src)
	return strings.Replace(src, findString, ``, 1)
}

// 删除正则匹配的内容（会把所有匹配的情况全删除）
func RemoveMatch(src, regex string) string {
	var re *regexp.Regexp

	re = regexp.MustCompile(regex)
	return re.ReplaceAllString(src, ``)
}

// 获取匹配到的组数
func MatchLen(text, regex string) int {
	var (
		re    *regexp.Regexp
		match [][]string
	)

	re = regexp.MustCompile(regex)
	match = re.FindAllStringSubmatch(text, -1)
	return len(match)
}

func MatchPreReplace(src, reg, repl string) (string, error) {
	compile, err := regexp2.Compile(reg, 0)
	if err != nil {
		return "", err
	}

	return compile.Replace(src, repl, 0, -1)
}

func MatchPreExist(str, reg string) (bool, error) {
	re := regexp2.MustCompile(reg, 0)
	matchString, err := re.MatchString(str)
	if err != nil {
		return false, err
	}
	return matchString, nil
}

func MatchPre(str, reg string) ([]string, error) {
	re, err := regexp2.Compile(reg, 0)
	if err != nil {
		return nil, err
	}

	var matches []string
	if m, _ := re.FindStringMatch(str); m != nil {
		for m != nil {
			matches = append(matches, m.String())
			m, _ = re.FindNextMatch(m)

			// log.Println(m.String()) // 奇怪，打印到最后一个竟然报空指针！
		}
		return matches, nil
	}

	return nil, nil
}
