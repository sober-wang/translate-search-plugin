/*
Packge common
公共处理包，读取文件，判断一些参数等
*/
package common

import (
	"unicode"
)

// IsOnlyEnglishLetters 是否只包含英文单词
func IsOnlyEnglishLetters(str string) bool {
	if len(str) == 0 {
		return false // 根据业务需求，空字符串可改为 true
	}
	for i := 0; i < len(str); i++ {
		c := str[i]
		if !((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
			return false
		}
	}
	return true
}

// IncludeMathOperator 判断输入字符是否包含数学运算符，货币单位
func IncludeMathOperator(str string) bool {
	exclude := map[rune]bool{
		'*': true,
		'(': true,
		')': true,
		'%': true,
	}

	//slog.Info("includ math operator", "str", str)
	for _, s := range str {
		if exclude[s] {
			return false
		}
		//slog.Info("for range ", "s", s)
		if unicode.IsSymbol(s) {
			return false
		}
	}
	return true

}
