 package main 

import (
	"flag"

	part "github.com/qydysky/part"
)

var (
	checkfile = flag.String("o", "", "output")
	root = flag.String("r", "", "root")
	checkDir = flag.String("i", "", "input")
	SplitString = flag.String("s", "\n", "SplitString")
	md5 = flag.Bool("m", true, "use md5")
)

func main() {
	flag.Parse()

	if *checkfile == "" {
		part.Logf().E("build {当存在以下必要参数时进行构建}")
		part.Logf().E("-i {你想要创建关于哪个目录的列表}")
		part.Logf().E("-o {列表输出的位置}")
		part.Logf().E("-r {列表的相对根路径 默认./ 可选}")
		part.Logf().E("-s {列表使用的分割符 默认\\n 可选}")
		part.Logf().E("-m {对每个文件计算md5 默认true 可选}")
		part.Logf().E("eg. go run main.go -i ./ -o t.txt")
		part.Logf().E("")
		part.Logf().E("check {当存在以下必要参数时进行检查}")
		part.Logf().E("-o {列表的位置}")
		part.Logf().E("-r {列表的相对根路径 默认./ 可选}")
		part.Logf().E("-s {列表使用的分割符 默认\\n 可选}")
		part.Logf().E("eg. go run main.go -o t.txt")
		return
	}

	if *checkDir != "" {
		part.Checkfile().Build(*checkfile,*root,*checkDir,*SplitString,*md5)
	}
	
	part.Checkfile().CheckList(*checkfile,*root,*SplitString)
}