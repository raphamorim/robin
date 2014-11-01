package libs

import (
	. "fmt"
)

func Help() {
	help := []string{
		"Usage: Robin <command>\n",
		"Description: A spotlight for command line",
		"More info: https://github.com/raphamorim/robin\n",
		"Commands:",
		" $ robin help      		[output usage information]",
		" $ robin version           	[output version number]",
		" $ robin domain <search>   	[output domainr search]"}

	Println()

	for index, element := range help {
		Println(element)
		index++
	}

	Println()
}

func Domain() {
	Println("Domain!!")
}

func Caniuse() {
	Println("Caniuse!!")
}
