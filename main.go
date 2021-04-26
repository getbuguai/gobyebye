package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
)

var (
	show bool
)

func init() {
	flag.BoolVar(&show, "show", false, "display deleted files .")
}

func main() {
	flag.Parse()

	log.SetFlags(0)
	if !show {
		log.SetOutput(ioutil.Discard)
	}

	fmt.Println("Hello Get BuGuai !!!")

	log.Println("----------------")
	err := Clean()
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	log.Println("----------------")
	log.Println("Clean over ...")
}
