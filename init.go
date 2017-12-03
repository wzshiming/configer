package configer

import (
	"encoding/json"
	"encoding/xml"

	toml "github.com/BurntSushi/toml"
	yaml "github.com/wzshiming/yaml"
	ini "gopkg.in/ini.v1"
)

var (
	jsonUnmarshal = json.Unmarshal
	xmlUnmarshal  = xml.Unmarshal
	yamlUnmarshal = yaml.Unmarshal
	tomlUnmarshal = toml.Unmarshal
	iniUnmarshal  = func(d []byte, v interface{}) error {
		f, err := ini.InsensitiveLoad(d)
		if err != nil {
			return err
		}
		return f.MapTo(v)
	}
)
