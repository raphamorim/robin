package libs

import (
	. "fmt"
	"./../utils"
	"net/http"
	"io/ioutil"
	"encoding/json"
	// "reflect"
)

func Help() {
	help := []string{
		"Usage: Robin <command>\n",
		"Description: A spotlight for command line",
		"More info: https://github.com/raphamorim/robin\n",
		"Commands:",
		" $ robin help      		[output usage information]",
		" $ robin version           	[output version number]",
		" $ robin domainr <search>   	[output domainr search]",
		" $ robin caniuse <search>   	[output caniuse search]"}

	Println()

	for index, element := range help {
		Println(element)
		index++
	}

	Println()
}

type Posts struct {
    userId int
    id  int
    title  string
    body  string
}

func Domain(search string) {
	res, err := http.Get("http://jsonplaceholder.typicode.com/posts")
	utils.Perror(err)

	Println(res)

	jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
    utils.Perror(err)

    // Println(reflect.TypeOf([]byte(jsonDataFromHttp))
    var domainrStruct []Posts

    err = json.Unmarshal([]byte(jsonDataFromHttp), &domainrStruct)
    utils.Perror(err)

    Println(domainrStruct) 
}

func Caniuse(search string) {
	Println("Searching for " + search)
}
