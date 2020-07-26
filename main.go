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

	if *checkfile == "" || *checkDir == "" {
		part.Logf().E("-i {checkDir} -o {checkfile} -s {SplitString}(default:\"\\n\")")
		part.Logf().E("eg. go run main.go -i ./ -o t.txt")
		return
	}

	part.Checkfile().Build(*checkfile,*checkDir,*SplitString)

	part.Checkfile().CheckList(*checkfile,*checkDir,*SplitString)
}