package main

import (
	"fmt"
	"github.com/sergi/go-diff/diffmatchpatch"
)

// go-diff

func main() {

	//var (
	//	text1 = "gocn vip"
	//	text2 = "goCN cool"
	//)
	//
	//d := diffmatchpatch.New()
	//
	//r := d.DiffMain(text1, text2, false)
	//
	//fmt.Println(r)
	//for _, k := range r {
	//	fmt.Printf("%v %q ", k.Type, k.Text)
	//}
	//fmt.Println()
	//fmt.Println(d.DiffPrettyHtml(r))
	//fmt.Println(d.DiffPrettyText(r))
	//

	originalText := "The quick brown fox\njumps over the lazy dog.\nwhat are you talking about?\nbi jia bi jia da."
	modifiedText := "The swift brown fox\nleaps over the lazy dog.\nwhat are you talking about?\nbii jjja jia da."

	// 创建 DiffMatchPatch 对象
	dmp := diffmatchpatch.New()

	// 将文本分割为基于行的数组，并获取字符到行的映射关系
	chars1, chars2, lineArray := dmp.DiffLinesToChars(originalText, modifiedText)

	// 对比两个字符数组，获取字符级别的差异
	charDiffs := dmp.DiffMain(chars1, chars2, false)

	// 我们现在有了字符级别的 diffs，但我们需要行级别的 diffs
	lineDiffs := dmp.DiffCharsToLines(charDiffs, lineArray)

	//fmt.Println(dmp.DiffPrettyHtml(lineDiffs))
	fmt.Println(dmp.DiffPrettyText(lineDiffs))
}
