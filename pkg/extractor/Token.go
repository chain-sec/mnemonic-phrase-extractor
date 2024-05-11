package extractor

import "unicode/utf8"

// 表示字符串里
type Token struct {

	// 用于表示字符串开始的位置
	Pos uint

	// 表示这个Token本身的内容，通常是一个单词
	Content string
}

// IsWord 判断此Token是否是一个单词
func (x *Token) IsWord() bool {
	for i:=0; i<len(){

	}
}

// 返回结束的位置
func (x *Token) EndPos() uint {
	return x.Pos + uint(utf8.RuneCountInString(x.Content))
}

