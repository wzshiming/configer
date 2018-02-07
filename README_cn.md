# configer
Golang配置工具，支持INI，XML，YAML，JSON，HCL，TOML，Shell环境

 - [English](./README.md)
 - [简体中文](./README_cn.md)

## 安装
``` bash
go get -u -v gopkg.in/configer.v1
```

## 使用

[API Documentation](http://godoc.org/gopkg.in/ffmt.v1)

## 优先级 
env > default > conf

## 示例

[examples](./examples/main.go)

``` golang
package main

import (
	configer "gopkg.in/configer.v1"
	ffmt "gopkg.in/ffmt.v1"
)

func main() {
	examples1()
}

type BB struct {
	Hello   string `configer:"world"`            // 取默认值 "world"
	Shell   string `configer:",env" env:"SHELL"` // 从 env 环境变量里取
	EnvNone string `configer:",env" env:"NONE"`  // 空的 env
}

type TT struct {
	LoadFilePath string `configer:"./examples1.json,env"`      // 加载文件的路径
	BB           BB     `configer:",load" load:"LoadFilePath"` // 加载路径字段
}

func examples1() {
	b := BB{}

	configer.Load(&b)
	ffmt.Puts(b)
	/*
		{
		 Hello:   "world"
		 Shell:   "/bin/bash"
		 EnvNone: ""
		}
	*/

	configer.Load(&b, "./examples1.json")
	ffmt.Puts(b)
	/*
		{
		 Hello:   "json"
		 Shell:   "/bin/bash"
		 EnvNone: "env none"
		}
	*/

	t := TT{}
	configer.Load(&t)
	ffmt.Puts(t)
	/*
		{
		 LoadFilePath: "./examples1.json"
		 BB:           {
		  Hello:   "json"
		  Shell:   "/bin/bash"
		  EnvNone: "env none"
		 }
		}
	*/
}

```

文件: examples1.json:

``` json
{
    "Hello": "json",
    "Shell": "Priority default < env",
    "EnvNone": "env none"
}
```

## MIT许可证

版权所有©2017-2018 wzshiming <[https://github.com/wzshiming](https://github.com/wzshiming)>。

MIT是[MIT许可证](https://opensource.org/licenses/MIT)许可的开源软件。
