package main

import (
	"flag"
	"fmt"
	"os"
	"testing"
)

var flagvar int
var s string
var bFlag bool

func init() {
	flag.IntVar(&flagvar, "a", 1234, "help message for flagname")
	flag.StringVar(&s, "b", "hello", "help string")
	flag.BoolVar(&bFlag, "c", true, "xx")
}

func TestStartCmd(t *testing.T) {
	flag.Parse()
	fmt.Println(flagvar)
	fmt.Println(s)
	fmt.Println(bFlag)
	fmt.Println(flag.Args())
	fmt.Println(os.Args)
}
