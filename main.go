package main

import (
	"./libs"
	"./utils"
	"os"
	"reflect"
)

func Validate(input string) bool {
	commands := map[string]bool{
		"help":    true,
		"domainr":  true,
		"caniuse": true,
		"isitup": true,
	}

	if commands[input] {
		return true
	}

	return false
}

func Run(m map[string]interface{}, name string, params ...interface{}) (result []reflect.Value, err error) {
	f := reflect.ValueOf(m[name])
	if len(params) != f.Type().NumIn() {
		utils.ThrowMsg(2)
		return
	}
	in := make([]reflect.Value, len(params))
	for k, param := range params {
		in[k] = reflect.ValueOf(param)
	}
	result = f.Call(in)
	return
}

func Execute(task string, params string) {
	funcs := map[string]interface{}{
		"help":    libs.Help,
		"domainr":  libs.Domain,
		"caniuse": libs.Caniuse,
		"isitup": libs.Isitup,
	}

	Run(funcs, task, params)
}

func main() {
	args := os.Args

	if len(args) > 1 {
		if Validate(args[1]) {
			if len(args) > 2 {
				Execute(args[1], args[2])

			} else {
				Execute(args[1], "")
			}
		} else {
			utils.ThrowMsg(1)
		}
	} else {
		utils.ThrowMsg(0)
	}
}
