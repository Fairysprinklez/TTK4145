package cost

import "../config"

//NEW IDEA: so if the Lift-argument in calculateCost has ONLY requests that are verified, we can use Ander's
//algorithm to naively finding the optimal targetFloor and motorDir (Like TilDat) and solve race-conditions between
//different nodes by just choosing one where calculateCost is called. NB: calculateCost, can NOT be run concurrently

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

func requestsAbove(lift config.Lift) {
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
func calculateCost(lift config.Lift) (int, config.MotorDirection) {
	haveRequestsAbove, floor := requestsAbove(lift)
	haveRequestsBelow, floor := requestsBelow(lift)
	switch lift.MotorDir {
	case config.MD_Up:
		if haveRequestsAbove {
			return floor, config.MD_Up
		} else if haveRequestsBelow {
			return floor, config.MD_Down
		}
	case config.MD_Down:
		if haveRequestsBelow {
			return floor, config.MD_Down
		} else if haveRequestsAbove {
			return floor, config.MD_Up
		}
	case config.MD_Stop:
		if haveRequestsBelow {
			return floor, config.MD_Down
		} else if haveRequestsAbove {
			return floor, config.MD_Up
		}
	default:
		return -1, MD_Stop //Don't really know if we want to do this?
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
