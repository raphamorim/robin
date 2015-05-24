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
		" $ robin caniuse <search>   	[output caniuse search]",
		" $ robin is-up <url>   [Check whether a website is up or down]",}

	Println()

	for index, element := range help {
		Println(element)
		index++
	}

	Println()
}

type Posts struct {
    UserId int
    Id  int
    Title  string
    Body  string
}

type IsitupStruct struct {
    Domain  string
    Port  int
    Status_Code int
    Response_Ip string
    Response_Code int
    Response_Time float64
}

func requestJson(url string ) {
	
}

func Isitup(url string) {
	apiUrl := "http://isitup.org/" + url +  ".json"
	res, err := http.Get(apiUrl)
	utils.Perror(err)

	jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
    utils.Perror(err)

    var isitupStruct IsitupStruct
    err = json.Unmarshal(jsonDataFromHttp, &isitupStruct)
    utils.Perror(err)

    if isitupStruct.Status_Code == 3 {
    	Println("Invalid domain!")
    	return
    }

    if isitupStruct.Status_Code == 2 {
    	Println(url + " is Down!")
    	return
    }

    if isitupStruct.Status_Code == 1 {
    	Println(url + " is Up!")
    	return
    }
}

func Domain(search string) {
	res, err := http.Get("http://jsonplaceholder.typicode.com/posts")
	utils.Perror(err)

	jsonDataFromHttp, err := ioutil.ReadAll(res.Body)
    utils.Perror(err)

    // Println(reflect.TypeOf([]byte(jsonDataFromHttp))
    var domainrStruct []Posts
    err = json.Unmarshal([]byte(jsonDataFromHttp), &domainrStruct)
    utils.Perror(err)
}

func Caniuse(search string) {
	Println("Searching for " + search)
}
