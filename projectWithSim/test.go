package main

import (
	"fmt"
	"strings"
	"strconv"
)

func main() {
	test := "123.542.656.456"
	test = strings.Split(test, ".")
	lastChars := test[len(test)-3:]
	if idInt, err := strconv.Atoi(lastChars); err == nil {
  		fmt.Println("%v: %T", idInt, idInt)
    				//this should break if we can't find the idInt
	}
}