package main

import (
	"fmt"
    "os"
)

func validate(input string) bool {
    commands := map[string]bool {
        "domain"  : true,
        "caniuse" : true,
    }

    if commands[input] {
        return true
    }

    return false
}

func throwMsg(msg int) {
    messages := [...]string{
        "[Robin] Welcome Batman!\n\n version: v0.1.0\n author: Raphael Amorim\n ",
        "[Robin] Hey Batman, I've founded a error!\n\n Error: Invalid command\n Suggestion: use `robin help` to see all commands\n",
        "[Robin] Hey Batman, please sent a parameter!" }

    fmt.Println(messages[msg])
}

func execute(task string, params string) {
    fmt.Printf(" Executing...\n\n Task: %s\n Params: %s\n\n ", task, params)
}

func main() {
    args := os.Args
    argsLen := len(args)

    if argsLen > 1 {
        if validate(args[1]) {
            if argsLen > 2 {
                execute(args[1], args[2])
            } else {
                throwMsg(2)
            }
        } else {
            throwMsg(1)
        }
    } else {
        throwMsg(0)
    }
}

