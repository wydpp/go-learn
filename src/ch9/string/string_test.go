package string

import "testing"

func TestString(t *testing.T) {
	var s string
	t.Log(s)
	s = "Hello"
	t.Log(len(s))
	//s[1]  = 3 string是不可变的 byte slice
	s = "\xE4\xB8\xA5" //任何二进制数据
	t.Log(s)
	t.Log(len(s))
	s = "中"
	t.Log(len(s)) //byte数

	c := []rune(s)
	//rune代表unicode
	t.Log(len(c))
	t.Logf("中 unicode %x", c[0])
	t.Logf("中 UTF8 %x", s)

}

func TestStringToRune(t *testing.T) {
	s := "中华人民共和国"
	for _, c := range s {
		//[1]都对应第一个参数
		//%c,%d,%x
		//range 遍历转成了 rune?
		t.Logf("%[1]c %[1]d", c)
	}
}
