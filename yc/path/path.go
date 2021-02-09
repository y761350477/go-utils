package path

import (
	"github.com/chenhg5/collection"
	"github.com/y761350477/go-utils/yc/regex"
	"os"
	"path/filepath"
)

// @Summary 判断目录下是否包含指定文件
// @Param   path string "文件路径"
// @Param   filename string "文件名称"
// @Return  bool "是否存在"
// @Return  error "异常信息"
func IsExist(path string, filename string) (bool, error) {
	var (
		f     *os.File
		names []string
		err   error
	)

	f, err = os.Open(path)
	if err != nil {
		return false, err
	}
	names, err = f.Readdirnames(1)
	if err != nil {
		return false, err
	}
	return collection.Collect(names).Contains(filename), nil
}

// @Summary 是否是指定后缀的文件
// @Param   path string "文件路径"
// @Param   suffix string "后缀名称（包含 .）"
// @Return  error "异常信息"
func IsExt(path, suffix string) bool {
	var fileSuffix string
	fileSuffix = filepath.Ext(path)
	return suffix == fileSuffix
}

// @Summary 获取上级目录文件名
// @Param   path string "文件路径"
// @Param   suffix string "后缀名称（包含 .）"
// @Return  string "目录路径"
func GetParentDirName(path string) string {
	dir := filepath.Dir(path)
	return regex.RegexRemove(dir, `.*\\+`)
}
