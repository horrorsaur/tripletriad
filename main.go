package main

import (
	"flag"
)

var (
	boardSize = flag.Int("size", 3, "sets the triple triad boards size. the default is 3. the maxiumum is 5 (experimental)")
)

func main() {
	flag.Parse()
	flag.Usage()
}
