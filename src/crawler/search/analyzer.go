package search

import (
	"fmt"
	"regexp"
)

func Analysis(content string) {

	//正则匹配 css/js标签并去除
	//fmt.Println(content)
	re := regexp.MustCompile(`<title>*</title>`)
	matched := re.FindAllStringSubmatch(content, -1)
	fmt.Println(matched)
/*	for _, match := range matched {
		fmt.Printf("email is: %s, domain is: %s\n", match[0], match[1])
	}*/


}