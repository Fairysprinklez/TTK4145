package cost

import (
	"../config"
	"strconv"
)

func CalculateNextFloorForAllNodes(nodeMapInCh chan config.NodeMap, thisLiftOutCh chan config.Lift) {
	for {
		//setting all hallRequests true
		var validRequests [config.NumFloors][config.NumButtons]bool
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons-1; b++ {
				validRequests[f][b] = true //requests that are present in all nodes
			}
			//the cabRequests in every floor is initialized to false
			validRequests[f][config.NumButtons-1] = false
		}

		nodeMap := <- nodeMapInCh
		aliveNodes = 0

		for id := range nodeMap {
			liftNode := nodeMap[id]
			if liftNode.Alive {
				aliveNodes++
				//finding all hallRequests that are common for every alive liftNode
				for f := 0; f < config.NumFloors; f++ {
					for b := 0; b < config.NumButtons-1; b++ {
						validRequests[f][b] = (validRequests[f][b] && lift.Requests[f][b])
					}
				}
			}
		}

		sortedIdSlice := make([]int, aliveNodes)
		sortedNextFloorSlice := make([]int, aliveNodes)
		for id := range nodeMap {
			liftNode := nodeMap[id]
			if liftNode.Alive {
				id = 
				lastChars := id[len(id)-3:]
				if idInt, err := strconv.Atoi(lastChars); err == nil {
    				sortedIdSlice = append(sortedIdSlice, idInt)
    				//this should break if we can't find the idInt
				}
			}
		}

	}
}