package funcExample

import "os"

func GetFilesContent(fileNames []string) (res map[string]interface{}){
	res = map[string]interface{}{}
	for _, fileName := range fileNames {
		var (
			content []byte
			err error
		)
		if content, err = getFileContent(fileName); err != nil {
			res[fileName] = err
			continue
		}
		res[fileName] = string(content)
	}

	return res
}

func getFileContent (fn string) (c []byte, err error){
	f, err := os.Open(fn)

	// 不能在这里定义defer f.Close，因为打开文件可能失败，此时f是一个nil

	if err != nil {
		return c, err
	}

	defer f.Close()		// 使用了defer下面就可以少写一句f.Close()

	fileStat, _ := f.Stat()
	c = make([]byte, fileStat.Size())	// 指定 c 的字节长度为文件大小的长度，这样下面Read(c)的时候c才能接收到所有的文件内容，c的字节切片长度是多少就能接收到多少文件内容。

	_, err = f.Read(c)		// 用c接收文件字节流

	return c, err
}