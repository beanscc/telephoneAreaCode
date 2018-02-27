package telephoneAreaCode

import (
	"strings"
)

// AreaCode
func AreaCode(cityName string) (string, bool) {
	if 2 > len([]rune(cityName)) {
		return "", false
	}

	// 过滤结尾的 市， 自治区，自治州，盟
	filterKey := []string{"市", "盟", "自治区", "自治区", "自治州", "地区"}
	// 内容匹配
	for _, key := range filterKey {
		cityName = sufFilter(cityName, key)
	}

	for _, area := range AreaCodeSlice {
		if ok := strings.HasPrefix(area.Name, cityName); ok {
			return area.ID, true
		}
	}

	return "", false
}

// 字符串s 从尾部去处 k
func sufFilter(s, k string) string {
	if ok := strings.HasSuffix(s, k); ok {
		return strings.Replace(s, k, "", 1)
	}

	return s
}
