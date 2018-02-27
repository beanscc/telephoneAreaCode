# telephoneAreaCode
根据城市名称获取国内固定电话长途区号

> 数据来源于中华人民共和国工业和信息化部（http://www.miit.gov.cn/n1146285/n1146352/n3054355/n3057709/n3057714/c5622616/content.html)，也可通过电信网码号管理系统(http://miinac.gov.cn/)查询获取

## 使用方法

可参考test方法
```go
func Test_AreaCode(t *testing.T) {
	cites := []string{
		"北京市",
		"上海",
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
```