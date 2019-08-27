package main

import (
	"./search"
)

func main() {

	cont := search.FetchByUrl("http://api.goseek.cn/")

	search.Analysis(cont)


}
