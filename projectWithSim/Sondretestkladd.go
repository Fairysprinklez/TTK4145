package main

import (
	"./config"
	"fmt"
)

func main(){

	nodeMap := make(map[string]config.Lift)
	nodeMap["test"] = Lift{
		"test",
		true,
		-1, 
		0, 
		{{false, false, false},
		 {false, false, false},
		 {false, false, false},
		 {false, false, false}}}
	
	fmt.Printf(nodeMap["test"])
	


}