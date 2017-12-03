package configer

import "testing"

type td struct {
	A     string `configer:"test"`
	Path1 string `configer:",env" env:"GOPATH"`
	Path2 string `configer:"ptah,env"`
	Sub   [1]struct {
		C  int     `configer:"10"`
		C1 float64 `configer:"10.2"`
		C2 float64 `configer:"-"`
	}
}

func TestTag(t *testing.T) {
	b := [2]td{}
	ProcessTags(&b)
	t.Log(b)
}
