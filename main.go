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

	part.Checkfile().Build(*checkfile,*checkDir,*SplitString)

	part.Checkfile().Check(*checkfile,*checkDir,*SplitString)
}