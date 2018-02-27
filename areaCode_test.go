package telephoneAreaCode

import (
	"testing"
)

// go test -v -run Test_AreaCode
func Test_AreaCode(t *testing.T) {
	cites := []string{
		"西",
		"北京市",
		"上海",
		"神",
		"神农架",
	}

	for _, city := range cites {
		code, ok := AreaCode(city)
		if !ok {
			t.Errorf("not found. city=%q", city)
		}

		t.Logf("city=%q, code=%q", city, code)
	}
}
