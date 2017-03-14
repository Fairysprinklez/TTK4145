package cost

import "../config"
import "fmt"

//NEW IDEA: so if the Lift-argument in calculateCost has ONLY requests that are verified, we can use Ander's
//algorithm to naively finding the optimal targetFloor and motorDir (Like TilDat) and solve race-conditions between
//different nodes by just choosing one where calculateCost is called. NB: calculateCost, can NOT be run concurrently

//DISCLAIMER: ALL OF THIS IS CURRENTLY WIP
func requestsBelow(lift config.Lift) (bool, int) {
	for f := 0; f < lift.LastKnownFloor; f++ {
		for b := 0; b < config.NumButtons; b++ {
			if lift.Requests[f][b] {
				return true, f
			}
		}
	}
	return false, -1
}

func requestsAbove(lift config.Lift) (bool, int) {
	for f := lift.LastKnownFloor + 1; f < config.NumFloors; f++ {
		for b := 0; b < config.NumButtons; b++ {
			if lift.Requests[f][b] {
				return true, f
			}
		}
	}
	return false, -1
}

//returns the closest targetFloor and what motorDir for a given Lift, assumes verified requests matrix
func calculateCost(lift config.Lift, validRequest [config.NumFloors][config.NumButtons]bool) (int) {
	haveRequestsAbove, floor := requestsAbove(lift)
	haveRequestsBelow, floor := requestsBelow(lift)
	switch lift.MotorDir {
	case config.MD_Up:
		if haveRequestsAbove {
			return floor
		} else if haveRequestsBelow {
			return floor
		}
	case config.MD_Down:
		if haveRequestsBelow {
			return floor
		} else if haveRequestsAbove {
			return floor
		}
	case config.MD_Stop:
		if haveRequestsBelow {
			return floor
		} else if haveRequestsAbove {
			return floor
		}
	default:
		return -1 //Don't really know if we want to do this?
	}
}

func findMinimum(slice []int, keyA int, keyB int) (int) {
	if slice[keyA] < slice[keyB] {
		return keyA
	} else {
		return keyB
	}
}

func calculateCostForAllNodes(nodeMapInCh chan config.NodeMap, thisLiftOutCh chan config.Lift) {
	for {

		var validRequests [config.NumFloors][config.NumButtons]bool
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons; b++ {
				validRequests[f][b] = true //requests that are present in all nodes
			}
		}

		nodeMap := <- nodeMapInCh
		sortedIdSlice := make([]int, len(nodeMap))
		sortedNextFloorSlice := make([]int, len(nodeMap))
		for id := range nodeMap {
			lift = nodeMap[id]
			if lift.Alive {
				for f := 0; f < config.NumFloors; f++ {
					for b := 0; b < config.NumButtons-1; b++ {
						validRequests[f][b] = (validRequests[f][b] && lift.Requests[f][b])
					}
				}
				sortedNextFloorSlice = append(sortedNextFloorSlice, calculateCost(lift, validRequests))
				lastChars := id[len(id)-3:]
				if idInt, err := strconv.Atoi(lastChars); err == nil {
    				sortedIdSlice = append(sortedIdSlice, idInt)
    				//this should break if we can't find the idInt
				}

			}
		}
		for i := 0 ; i < len(sortedNextFloorSlice); i++ {
			for j := 0 ; j < len(sortedNextFloorSlice); j++ {
				if (sortedNextFloorSlice[i] == sortedNextFloorSlice[j]) && (i!=j) {
					minKey := findMinimum(sortedIdSlice, i, j)
					sortedNextFloorSlice[minKey] = -1
				}
			}
		}


	}
} 

/*//Version 1.0
//what kind of interface does this function have?
//it gets a map from NodeMapCompiler, but it also should update the ThisLift struct from main so that the
//FSM can do its magic. Or do you want something to bridge that?
//It could send a NodeMap both ways and main will extract the ThisLift struct? - Martin
//https://github.com/mortenfyhn/TTK4145-Heis/blob/master/Lift/src/queue/cost.go
//^link for "inspiration"^
func calculateCost(nodeMapInCh chan config.NodeMap, nodeMapOutCh chan config.NodeMap) {
	//TODO: this is propably important...
	//set Lift.Targetfloor to what it calculates as the optimal solution
	//can this be so "naive" that it just finds the next optimal floor for all nodes?
	for {
		var requests [config.NumFloors][config.NumButtons]bool
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons; b++ {
				validRequests[f][b] = true //requests that are present in all nodes
			}
		}
		nodeMap := <-nodeMapInCh
		for id := range nodeMap {
			lift = nodeMap[id]
			for f := 0; f < config.NumFloors; f++ {
				for b := 0; b < config.NumButtons-1; b++ {
					validRequests[f][b] = (validRequests[f][b] && lift.Requests[f][b])
				}
			}
		}


		nodeMapOutCh <- nodeMap
	}
}*/
