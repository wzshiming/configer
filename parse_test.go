package configer

import (
	"testing"
)

var parseData = []string{
	`ab: ss`,
	`{"AB": "ss"}`,
	`<xml><AB>ss</AB></xml>`,
	`AB = "ss"`,
	`AB = ss`,
}

func TestParse(t *testing.T) {
	for _, data := range parseData {
		var v struct {
			AB string
		}
		err := Parse([]byte(data), "", &v)
		if err != nil {
			t.Error(err)
		}
		if v.AB == "" {
			t.Error("Parse error")
		}
	}
}
