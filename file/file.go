package file

import (
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
)

//	获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

//	获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

//	检查文件存在与否
func CheckNotExist(src string) bool {
	_, err := os.Stat(src)

	return os.IsNotExist(err)
}

//	检查文件权限
func CheckPermission(src string) bool {
	_, err := os.Stat(src)

	return os.IsPermission(err)
}

//	文件不存在则创建文件
func IsNotExistMKDir(src string) error {
	if notExist := CheckNotExist(src); notExist == true {
		if err := MKDir(src); err != nil {
			return err
		}
	}

	return nil
}

//	创建文件
func MKDir(src string) error {
	err := os.MkdirAll(src, os.ModePerm)
	if err != nil {
		return err
	}

	return nil
}

//	打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
