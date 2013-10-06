package main

import (
	"flag"
	"fmt"
	"os"
)

var (
	versionflag = flag.Bool("v", false, "print version and exit")
	fileflag    = flag.String("f", "", "file to parse")
)

func main() {
	flag.Parse()

	if *versionflag {
		fmt.Println(Version())
		os.Exit(0)
	}

	if *fileflag == "" {
		fmt.Println("no file to parse, exiting...")
		os.Exit(0)
	}

	_, err := Parse(*fileflag)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		os.Exit(1)
	}
}
