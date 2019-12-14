package main

import (
	"flag"
	"fmt"
	"os"

	"github.com/akihiro-fukuchi/cnasb/pkg/stringutil"
)

func main() {
	flag.Parse()
	args := flag.Args()
	if len(args) != 1 {
		fmt.Println("error: expected 'grev KEYWORD'. KEYWORD is a required argument for the grev command")
		os.Exit(1)
	}
	fmt.Println(stringutil.Reverse(args[0]))
}
