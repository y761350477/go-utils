package yc_io

import (
	"bufio"
	"github.com/y761350477/go-utils/yc_regex"
	"github.com/y761350477/go-utils/yc_str"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"strings"
)

var (
	// Windows 文件命名中不合法字符
	ErrChars = []string{`<`, `>`, `:`, `*`, `?`, `"`, `|`, `//`}
)

// @Summary 读取文件内容（整体读取）
// @Param   filepath string "文件路径"
// @Return  string "文件内容"
// @Return  error "异常信息"
func ReadFile(filePath string) (string, error) {
	var (
		buf []byte
		err error
	)
	buf, err = ioutil.ReadFile(filePath)
	if err != nil {
		return "", err
	}
	return string(buf), nil
}

// @Summary 读取文件内容（逐行读取）
// @Param   filepath string "文件路径"
// @Return  string "文件内容"
// @Return  error "异常信息"
func ReadFileLine(filePath string) (string, error) {
	var (
		f    *os.File
		r    *bufio.Reader
		line []byte
		str  strings.Builder
		err  error
	)

	f, err = os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer f.Close()

	r = bufio.NewReader(f)
	for {
		line, _, err = r.ReadLine()
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		str.WriteString(string(line) + "\n")
	}

	return str.String(), nil
}

// @Summary 读取文件内容（指定缓存大小，默认 1024）
// @Param   filepath string "文件路径"
// @Param   bufferSize int "缓存大小，默认 1024"
// @Return  string "文件内容"
// @Return  error "异常信息"
func ReadFileBuffer(filePath string, bufSize ...int) (string, error) {
	var (
		file     *os.File
		reader   *bufio.Reader
		err      error
		buf      []byte
		n        int
		bufSizeT int
	)
	if len(bufSize) > 0 {
		bufSizeT = bufSize[0]
	} else {
		bufSizeT = 1024
	}

	file, err = os.Open(filePath)
	if err != nil {
		return "", err
	}
	defer file.Close()
	reader = bufio.NewReader(file)

	buf = make([]byte, bufSizeT)
	var str strings.Builder
	for {
		n, err = reader.Read(buf)
		if err == io.EOF {
			break
		}
		if err != nil {
			return "", err
		}
		str.Write(buf[:n])
	}
	return str.String(), nil
}

// @Summary 写入文件内容（清空）
// @Param   outFile string "文件路径"
// @Param   content string "内容"
// @Return  error "异常信息"
func WriterFile(outFile string, content string) error {
	var (
		f   *os.File
		err error
	)

	err = CheckFolderPath(outFile)
	if err != nil {
		return err
	}

	f, err = os.Create(outFile)
	if err != nil {
		return err
	}
	return writerFile(f, content)
}

// @Summary 写入文件内容（追加）
// @Param   outFile string "文件路径"
// @Param   content string "内容"
// @Return  error "异常信息"
func WriterFileAppend(outFile string, content string) error {
	var (
		f   *os.File
		err error
	)

	err = CheckFolderPath(outFile)
	if err != nil {
		return err
	}

	f, err = os.OpenFile(outFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0644)
	if err != nil {
		return err
	}
	return writerFile(f, content)
}

// @Summary 写入文件内容（追加）
// @Param   path string "文件路径"
// @Return  string "校验后的名称"
func CheckPathName(path string) string {
	var base = path
	for _, v := range ErrChars {
		base = strings.ReplaceAll(base, v, "")
	}
	return base
}

// @Summary 检测目录是否存在，不存在则创建
// @Param   path string "目录路径"
// @Return  error "异常信息"
func CheckFolderPath(path string) error {
	var (
		dir string
		err error
	)

	dir = filepath.Dir(path)
	err = MkdirFolder(dir)
	if err != nil {
		return err
	}
	return nil
}

// @Summary 写入文件内容（追加避免重复）
// @Param   outFile string "文件路径"
// @Param   content string "内容"
// @Return  error "异常信息"
func WriterFileAppendNoRepeat(outFile string, content string) error {
	var readSource string
	readSource, _ = ReadFileBuffer(outFile)
	if !strings.Contains(readSource, content) {
		return WriterFileAppend(outFile, content)
	}
	return nil
}

func writerFile(f *os.File, content string) error {
	defer f.Close()
	writer := bufio.NewWriter(f)
	writer.WriteString(content)
	writer.Flush()
	return nil
}

// @Summary 复制文件
// @Param   srcPath string "源文件路径"
// @Param   dstPath string "目标文件路径"
// @Return  error "异常信息"
func CopyFile(srcPath, dstPath string) error {
	var (
		src *os.File
		dst *os.File
		buf []byte
		err error
	)

	src, err = os.Open(srcPath)
	if err != nil {
		return err
	}
	defer src.Close()
	dst, err = os.Create(dstPath)
	if err != nil {
		return err
	}
	defer dst.Close()
	buf = make([]byte, 8192)
	_, err = io.CopyBuffer(dst, src, buf)
	if err != nil {
		return err
	}
	return nil
}

// @Summary 删除文件
// @Param   path string "文件路径"
// @Return  error "异常信息"
func DeleteFile(path string) error {
	var err error
	if err = os.Remove(path); err != nil {
		return err
	}
	return nil
}

// @Summary 判断文件夹是否存在, 不存在则创建
// @Param   path string "文件路径"
// @Return  error "异常信息"
func MkdirFolder(path string) error {
	var (
		exists bool
		err    error
	)

	exists, err = PathExists(path)
	if err != nil {
		return err
	}
	if !exists {
		os.MkdirAll(path, os.ModePerm)
	}
	return nil
}

// @Summary 判断文件 \ 文件夹是否存在
// @Param   path string "文件路径"
// @Return  error "异常信息"
func PathExists(path string) (bool, error) {
	var err error

	_, err = os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

// @Summary 获取指定目录下的所有文件，包含子目录下的文件
// @Param   path string "文件路径"
// @Param   filter string "过滤类型，非必填"
// @Return  files "文件列表"
// @Return  error "异常信息"
func GetAllFiles(dirPth string, filter ...string) ([]string, error) {
	var (
		files   []string
		dirs    []string
		filterT string
		err     error
	)

	if len(filter) > 0 {
		filterT = filter[0]
	}

	dir, err := ioutil.ReadDir(dirPth)
	if err != nil {
		return nil, err
	}

	PthSep := string(os.PathSeparator)
	for _, fi := range dir {
		if fi.IsDir() { // 目录, 递归遍历
			dirs = append(dirs, dirPth+PthSep+fi.Name())
			GetAllFiles(dirPth+PthSep+fi.Name(), filterT)
		} else {
			// 过滤指定格式
			if filterT != "" {
				ok := strings.HasSuffix(fi.Name(), filterT)
				if ok {
					files = append(files, dirPth+PthSep+fi.Name())
				}
			} else {
				files = append(files, dirPth+PthSep+fi.Name())
			}

		}
	}

	// 读取子目录下文件
	for _, table := range dirs {
		temp, _ := GetAllFiles(table, filterT)
		for _, temp1 := range temp {
			files = append(files, temp1)
		}
	}

	return files, nil
}

// @Summary 获取路径下的目录
// @Param   path string "文件路径"
// @Return  dirs "目录列表"
// @Return  error "异常信息"
func GetAllDirs(path string) ([]string, error) {
	var (
		dir      []os.FileInfo
		dirs     []string
		tempDirs []string
		err      error
	)

	dir, err = ioutil.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, fi := range dir {
		if fi.IsDir() {
			dirPath1 := path + yc_str.PthSep + fi.Name()
			tempDirs = append(tempDirs, dirPath1)
			dirs = append(dirs, dirPath1)
			GetAllDirs(dirPath1)
		}
	}

	// 读取子目录下文件
	for _, d := range tempDirs {
		temp, _ := GetAllDirs(d)
		for _, temp1 := range temp {
			dirs = append(dirs, temp1)
		}
	}
	return dirs, nil
}

// @Summary 删除目录下所有空文件夹（递归遍历内部）
// @Param   path string "文件路径"
// @Return  error "异常信息"
func DeleteEmptyDirs(path string) error {
	var (
		dirs []string
		err  error
	)

	dirs, err = GetAllDirs(path)
	if err != nil {
		return err
	}
	sort.StringSlice(dirs).Sort()
	for i := len(dirs) - 1; i >= 0; i-- {
		dir, err := ioutil.ReadDir(dirs[i])
		if err != nil {
			return err
		}

		dirSize := len(dir)
		if dirSize <= 0 {
			os.Remove(dirs[i])
		}
	}

	return nil
}

// @Summary 判断文件是否包含指定内容
// @Param 	path string "文件路径"
// @Param 	content string "内容（支持正则）"
// @Return 	bool "是否存在"
// @Return 	error "异常信息"
func IsContainsContent(path, content string) (bool, error) {
	buffer, err := ReadFileBuffer(path)
	if err != nil {
		return false, err
	}
	return yc_regex.MatchExist(buffer, content), nil
}
