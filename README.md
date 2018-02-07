# configer
Golang Configuration tool that support INI, XML, YAML, JSON, HCL, TOML, Shell Environment

## Install
``` bash
go get -u -v gopkg.in/configer.v1
```

## Usage

``` golang
// Config load
func Load(v interface{}, cfgpaths ...string) error // Load data and process tag data
func ProcessTags(config interface{}) error // Process tag data
```

## Priority 
env > default > conf

## Examples

[examples](./examples/main.go)

``` golang
b := struct {
    Hello   string `configer:"world"`            // Take the default value "world"
    Shell   string `configer:",env" env:"SHELL"` // Take the value of env
    EnvNone string `configer:",env" env:"NONE"`  // An empty env
}{}

configer.Load(&b)
/*
{
 Hello:   "world"
 Shell:   "/bin/bash"
 EnvNone: ""
}
*/

configer.Load(&b, "./examples1.json") // or configer.Load(&b, "./examples1.yml") 
/*
{
 Hello:   "json"
 Shell:   "/bin/bash"
 EnvNone: "env node"
}
*/
```

With examples1.json:

``` json
{
    "Hello": "json",
    "Shell": "Priority default < env",
    "EnvNone": "env none"
}
```

## MIT License

Copyright © 2017-2018 wzshiming<[https://github.com/wzshiming](https://github.com/wzshiming)>.

MIT is open-sourced software licensed under the [MIT License](https://opensource.org/licenses/MIT).

