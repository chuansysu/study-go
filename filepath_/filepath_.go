package filepath_

import (
	"fmt"
	"os"
	"path/filepath"
)

func DemoFilePath() {
	s := "F:\\golang_workspace\\leetcode\\aa.js"
	fmt.Println("path:", s)
	s = filepath.ToSlash(s) // 将 path 中平台相关的路径分隔符转换为 '/'
	fmt.Println("ToSlash:", s)
	s = filepath.FromSlash(s) // 将 path 中的 '/' 转换为系统相关的路径分隔符
	fmt.Println("FromSlash:", s)

	s1 := filepath.Dir(s)
	fmt.Println("Dir:", s1) // 获取path中最后一个分隔符之前的部分(不包含分隔符)
	s2 := filepath.Base(s)  // 获取path中最后一个分隔符之后的部分(不包含分隔符)
	fmt.Println("Base:", s2)
	s = filepath.ToSlash(s)
	s1, s2 = filepath.Split(s) // 获取path中最后一个分隔符前后的两部分
	fmt.Println("Split:", s1, s2)

	s1 = filepath.Ext(s) // 获取路径字符串中的文件扩展名
	fmt.Println("Ext:", s1)

	s1 = "/golang_workspace/leetcode/aa.js"
	s2 = "/golang_workspace/"
	s1, _ = filepath.Rel(s2, s1) //获取targpath相对于basepath的路径
	// 要求targpaht和basepath必须"都是相对路径"或都是"绝对路径"
	fmt.Println("Rel:", s1)

	s1 = "golang_workspace"
	s2 = "leetcode/aa.js"
	s1 = filepath.Join(s1, s2) // 将elem中的多个元素合并成一个路径，忽略空元素，清理多余字符
	fmt.Println("Join:", s1)

	//s1 = filepath.Clean("/.../..../abc/abc") // 清除path中多余的字符
	s1 = filepath.Clean("C:/a/b/../c") // 清除path中多余的字符
	fmt.Println("Clean:", s1)

	s1 = "C:/home/gopher"
	s2 = ".bashrc"
	f1 := filepath.IsAbs(s1) // 判断该路径是否是绝对路径
	f2 := filepath.IsAbs(s2)
	fmt.Println("IsAbs:", f1, f2) // true false

	s2, _ = filepath.Abs(s2)
	fmt.Println("Abs:", s2)

	s1 = "/a/b/c;/usr/bin"
	sList := filepath.SplitList(s1) // 按os.PathListSeparator即(;)将路径进行分割
	fmt.Println("SplitList:", sList)

	fmt.Println("VolumeName:", filepath.VolumeName(s))

	r := "/home/catch/*"
	s = "/home/catch/foo"
	f, _ := filepath.Match(r, s)
	fmt.Println("Match: ", f)

	r = "C:\\Users\\Administrator\\[C]*"
	sList, _ = filepath.Glob(r)
	fmt.Println("Glob:", sList)
	fmt.Println(os.Getwd())

	pwd, _ := os.Getwd()
	filepath.Walk(pwd, func(fpath string, info os.FileInfo, err error) error {
		if match, err := filepath.Match("???", filepath.Base(fpath)); match {
			fmt.Println("Walk path:", fpath)
			fmt.Println("Walk info:", info)
			return err
		}
		return nil
	})

	var i interface{} = 3
	fmt.Print(i)
}
