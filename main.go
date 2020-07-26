 package main 

import (
	"flag"

	part "github.com/qydysky/part"
)

var (
	checkfile = flag.String("o", "", "output")
	checkDir = flag.String("i", "", "input")
	SplitString = flag.String("s", "\n", "SplitString")
)

func main() {
	flag.Parse()

	if *checkfile == "" {
		part.Logf().E("build:")
		part.Logf().E("-i {checkDir} -o {checkfile} -s {SplitString}(default:\"\\n\")")
		part.Logf().E("eg. go run main.go -i ./ -o t.txt")
		part.Logf().E("")
		part.Logf().E("check:")
		part.Logf().E("-o {checkfile} -s {SplitString}(default:\"\\n\")")
		part.Logf().E("eg. go run main.go -o t.txt")
		return
	}

	if *checkDir != "" {
		part.Checkfile().Build(*checkfile,*checkDir,*SplitString)
	}
	
	part.Checkfile().CheckList(*checkfile,*SplitString)
}