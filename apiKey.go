package main

import (
	"fmt"
	"io/ioutil"
)

func readKey(filename string) string {
	dat, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Printf("Error in reading api key: %v", err)
	}

	return string(dat)
}


