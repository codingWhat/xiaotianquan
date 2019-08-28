package main

import (
	"crawler/search"
	"fmt"
)

func main() {


	// 待爬取网页链接: links.bin
	//linksFilePrefix := "links"

	// 网页判重文件: bloom_filter.bin
	//filterFilePrefix := "bloom_filter.bin"

	// 原始网页存储文件: doc_raw.bin
	//rawFilePrefix := "doc_raw"


	//res, _ := fileExists(getIndexesPath() + "links.bin")

	//return
	cont := search.FetchByUrl("")
	fmt.Println(cont)

	search.Analysis(cont)

	return
	//linksFilePrefix := "links"

	//file_push_conents(getIndexesPath() + makeFileName(linksFilePrefix), "This is a test");*/
}

/*
func file_push_conents(relativePath, content string)  {
	file, err := os.Create(relativePath)
	if err != nil {
		panic(err)
	}

	_, _ = file.Write([]byte(content))

}

func makeFileName(prefix string) string  {
	filePrefix := ".bin"
	date := time.Now().Format("2006-01-02") //time.Now().Format("2006-01-02 15:04:05")
	return prefix + date + filePrefix
}

func getIndexesPath() string {
	basePath, _ := GetBasePath()
	return fmt.Sprintf("%s%s%s%s", basePath, string(os.PathSeparator), "indexes", string(os.PathSeparator))
}

func getRawDocSavePath() string {
	basePath, _ := GetBasePath()

	return fmt.Sprintf("%s%s%s%s", basePath, string(os.PathSeparator), "raw", string(os.PathSeparator));
}

func fileExists(path string) (bool, error) {
	_, err := os.Stat(path)

	if err == nil {
		return true, nil
	}

	if os.IsNotExist(err) {
		return false, nil
	}

	return  false, nil
}

func GetBasePath()(string, error) {
	return filepath.Abs(filepath.Dir(os.Args[0]))
}*/