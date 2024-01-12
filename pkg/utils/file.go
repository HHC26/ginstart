package utils

import "os"

// CheckNotExist check if the file exists
func CheckFileNotExist(src string) bool {
	_, err := os.Stat(src)
	return os.IsNotExist(err)
}

func CreateFile(path string) error {
	fp, err := os.Create(path) // 如果文件已存在，会将文件清空。
	if err != nil {
		return err
	}
	defer fp.Close() //关闭文件，释放资源。
	return nil
}
