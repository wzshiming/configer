package main

import (
	configer "gopkg.in/configer.v1"
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	examples1()
}

func examples1() {
	b := struct {
		Hello   string `configer:"world"`            // Take the default value "world"
		Shell   string `configer:",env" env:"SHELL"` // Take the value of env
		EnvNone string `configer:",env" env:"NONE"`  // An empty env
	}{}

	configer.Load(&b)
	ffmt.Puts(b)

	configer.Load(&b, "./examples1.json")
	ffmt.Puts(b)
}
