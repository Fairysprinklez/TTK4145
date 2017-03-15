package main

import (
	//"./config"
	//"./network/localip"
	"fmt"
	//"strings"
	//"strconv"
)
/*func initializeLiftData() config.Lift {
	var lift config.Lift
	id, err := localip.LocalIP()
	if err == nil {
		lift = config.Lift{
			ID: id,
			Alive: true,
			LastKnownFloor: -1,
			TargetFloor: -1,
			MotorDir: config.MD_Stop,
			Behaviour: config.LiftIdle}

	}

	return lift
}*/

var test []int
//var test2 []int
//var lifts []config.Lift
func main() {

	test = append(test,2)
	fmt.Println(test)
	test = nil
	fmt.Println(test)
	test = append(test,1)
	/*test = append(test, 2)
	test2 = append(test2, 3)
	test2 = append(test2, 4)
	test = append(test2..., test)*/

	fmt.Println(test)

	/*lifts = make([]config.Lift, 0)

	testlift := initializeLiftData()
	testlift2 := initializeLiftData()
	testlift2.ID = "hei"

	lifts = append(lifts, testlift)
	lifts = append(lifts, testlift2)

	fmt.Println(lifts)
	lifts = lifts[:len(lifts)-1]
	fmt.Println(lifts)

	/test := "123.542.656.456"
	tmp := strings.Split(test, ".")
	if idInt, err := strconv.Atoi(tmp[len(tmp)-1]); err == nil {
  		fmt.Println("%v: %T", idInt, idInt)
    				//this should break if we can't find the idInt
	}*/
}