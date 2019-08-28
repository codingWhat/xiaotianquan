package search

import (
	"io/ioutil"
	"log"
	"net/http"
)

func FetchByUrl(url string) string  {

	//resp, err := http.Get("http://api.goseek.cn/")
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	result, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		panic(err);
	}

/*
   	buf := make([]byte, 4096);
   	//var  result string
for   {
		num , err := resp.Body.Read(buf)

		if err != nil && err != io.EOF {
			panic(err)
		}

		if num == 0{
			break
		}

		result += (string)(buf[:num])
	}*/

	return  string(result)
}
