package main

import (
	"fmt"
	"strings"
)

// EqualFold(s,t string)bool				检查两个字符串是否相等,****忽略大小写****
// TrimSpace(s string)string				去掉首尾空白符
// HasPrefix(s string, prefix string)bool	是否有前缀
// HasSuffix(s string, suffix string)bool	是否有后缀
// Index(s string, str string)int			判断str在s中**首次**出现的位置，如果没有出现，则返回-1
// LastIndex(s string, str string)int		判断str在s中**最后**出现的位置，如果没有出现，则返回-1
// IndexAny(s, chars string)int				返回包含chars中任意字符的位置
// IndexByte(s string, c byte)int			返回字符串中第一个字符实例的索引。如果找到，则返回以0开头的特定位置。如果找不到，则返回-1
// Replace(s string,old string,new string,n int)	n指定要在字符串中替换的字符数。如果n小于0，则替换次数数没有限制
func main() {
	fmt.Println(strings.EqualFold("TopGoer", "TOPGOER"))       // true
	fmt.Println(strings.TrimSpace(" Topgoer    "))             // Topgoer
	fmt.Println(strings.HasPrefix("Topgoer", "Top"))           // true
	fmt.Println(strings.HasSuffix("Topgoer", "er"))            // true
	fmt.Println(strings.Index("Topgoer", "goer"))              // 3
	fmt.Println(strings.LastIndex("hellollolloworld!", "llo")) // 8
	fmt.Println(strings.IndexAny("hellollolloworld!", "d!lo")) // 2 -> l

	// IndexByte
	fmt.Println(strings.IndexByte("Topgoer", 'g')) // 3
	var u byte
	u = 1
	fmt.Println(strings.IndexByte("5221-topgoer", u)) // -1
	u = '1'
	fmt.Println(strings.IndexByte("5221-topgoer", u)) // 3

	// Replace
	// Australia Japan Canada Indiana
	// Australia Japanese Canada Indiana
	// Australia Japanese Caneseada Indiana
	// Australia Japanese Caneseada Indianesea
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 0))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 1))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", 2))
	fmt.Println(strings.Replace("Australia Japan Canada Indiana", "an", "anese", -1))

	// Title
	fmt.Println(strings.Title("i like golang")) // I Like Golang

	fmt.Println(strings.TrimLeft("00012300", "0"))   // 12300
	fmt.Println(strings.TrimPrefix("00012300", "0")) // 0012300
}
