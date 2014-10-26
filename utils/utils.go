package utils

import (
	. "fmt"
)

func ThrowMsg(msg int) {
	/* Availables Messages:
	   0 - Default Msg
	   1 - Invalid Command Error
	   2 - Missing Parameter
	*/

	messages := [...]string{
		"Welcome Batman!\n\n version: v0.1.0\n author: Raphael Amorim\n ",
		"Hey Batman, I've founded a error!\n\n Error: Invalid command\n Suggestion: use `robin help` to see all commands\n",
		"Hey Batman, please sent a valid parameter!"}

	Println("[Robin]", messages[msg])
}
