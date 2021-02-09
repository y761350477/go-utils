package archive

import "github.com/mholt/archiver"

// @Summary 判断 RAR 文件是否损坏
// @Param   path string "文件路径"
// @Return  bool "损毁为 true"
// @Return  error "异常信息"
func CheckRarFile(path string) (bool, error) {
	var (
		r   *archiver.Rar
		err error
	)

	r = archiver.NewRar()
	err = r.Walk(path, func(f archiver.File) error {
		return nil
	})
	if err != nil {
		return true, err
	}
	return false, nil
}
