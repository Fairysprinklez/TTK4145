package cost

import (
	"../config"
	"strconv"
	"math"
	"sort"
)
func (slice []config.Lift) Len() int{
	return len(slice)
}

func (slice []config.Lift) Less(i, j int) bool{
	return slice[i].ID < slice[j].ID
}

func (slice []config.Lift) Swap(i, j int){
	slice[i], slice[j] = slice[j], slice[i]
}


func requestsBelow(lift config.Lift, validHallRequests [config.NumFloors][config.NumButtons-1]bool) ([]int) {
	var f int
	var floors []int
	for f = lift.LastKnownFloor - 1; f >=0; f-- {
		if validHallRequests[f][config.B_HallDown] || lift.Requests[f][config.B_Cab] {
			floors = append(floors, f)
		}
	}
	for f = 0; f < config.NumFloors; f++ {
		if validHallRequests[f][config.B_HallUp] || lift.Requests[f][config.B_Cab] {
			floors = append(floors, f)
		}
	}
	for f = config.NumFloors-1; f >= lift.LastKnownFloor; f++ {
		if validHallRequests[f][config.B_HallDown]{
			floors = append(floors, f)
		}
	}

	return floors
}

func requestsAbove(lift config.Lift, validHallRequests [config.NumFloors][config.NumButtons-1]bool) ([]int) {
	var f int
	var floors []int
	for f = lift.LastKnownFloor + 1; f < config.NumFloors; f++ {
		if (validHallRequests[f][config.B_HallUp]	|| lift.Requests[f][config.B_Cab]){
			floors = append(floors, f)
		}	
	}
	for f = config.NumFloors-1; f >= 0; f-- {
		if (validHallRequests[f][config.B_HallDown] || lift.Requests[f][config.B_Cab]){
			floors = append(floors, f)
		}
	}
	for f = 0; f <= lift.LastKnownFloor; f++{
		if (validHallRequests[f][config.B_HallUp]){
			floors = append(floors, f)
		}
	}

	return floors
}

/*func calculateTargetFloor(lift config.Lifts, orderedRequests [4]int) (int){
	for i := 0; i =< 4; i++{
		if orderedRequests[i] != -1 {
			if lift.Requests[orderedRequests[i]][config.B_Cab]{
				return orderedRequests[i]
			} 
		}else{
			return -1
		}
	}*/

}

//returns the closest targetFloor and what motorDir for a given Lift, assumes verified requests matrix
func prioritiseOrders(lift config.Lift, validHallRequests [config.NumFloors][config.NumButtons-1]bool) ([]int) {
	var requests []int
	requestsAbove := requestsAbove(lift, validHallRequests)
	requestsBelow := requestsBelow(lift, validHallRequests)
	switch lift.MotorDir {
	case config.MD_Up:
		requests = requestsAbove
		requests = append(requests, requestsBelow...)
	case config.MD_Down:
		requests = requestsBelow
		requests = append(requests, requests)
	case config.MD_Stop:
		distanceUp := abs(lift.LastKnownFloor - requestsAbove[0])
		distanceDown := abs(lift.LastKnownFloor - requestsBelow[0])
		if distanceUp < distanceDown{
			requests = requestsAbove
			requests = append(requests, requestsBelow...)
		}else {
			requests = requestsBelow
			requests = append(requests, requests)
	}
	return requests
}

/*func sortLiftSlice(lifts []config.Lift) ([]config.Lift) {
	var sortedSlice []config.Lift
	for i := 0; i < len(lifts); i++ {
		tmp := lifts[i].ID
		tmp = strings.Split(tmp, ".")
		if idInt, err := strconv.Atoi(tmp[len(tmp)-1]); err == nil {
			if len(sortedSlice) == 0{
				sortedSlice = append(sortedSlice, lifts[i])
			}else{
				for j := 0; j < sortedSlice; j++{

				}
			}
		}

	}
}*/

var lifts []config.Lift
var validHallRequests [config.NumFloors][config.NumButtons-1]bool

func CalculateNextFloorForAllNodes(nodeMapInCh chan config.NodeMap, thisLiftOutCh chan config.Lift, doneReadingMap chan bool, myID string) {
	lifts = make([]config.Lift)
	for {
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons-1; b++ {
				validHallRequests[f][b] = true //requests that are present in all nodes
			}
		}

		nodeMap := <- nodeMapInCh
		aliveNodes = 0
		for id := range nodeMap {
			liftNode := nodeMap[id]
			if liftNode.Alive {
				lifts = append(lifts, liftNode)
				aliveNodes++
				for f := 0; f < config.NumFloors; f++ {
					for b := 0; b < config.NumButtons-1; b++ {
						validHallRequests[f][b] = (validHallRequests[f][b] && lift.Requests[f][b])
					}
				}
			}
		}
		lightMatrix := validHallRequests
		doneReadingMap <- true
		sort.Sort(lifts)

		//new block
		for i := 0; i < len(lifts); i++ {
			requests := prioritiseOrders(lifts[i], validHallRequests)
			target := requests[1]
			lifts[i].targetFloor = target
			if lifts[i].ID == myID{
				for f := 0; f < config.NumFloors; f++{
					lifts[i].Requests[f][B_HallUp] = lightMatrix[f][B_HallUp]
					lifts[i].Requests[f][B_HallDown] = lightMatrix[f][B_HallDown]
				}
				thisLiftOutCh <- lifts[i]
			}
			validHallRequests[target][config.B_HallUp] = false
			validHallRequests[target][config.B_HallDown] = false
		}












		//old block (unfinished)
		/*var orderedRequests [len(lifts)][4]int
		for i := 0; i <len(lifts); i++{
			orders := prioritiseOrders(lifts[i], validHallRequests)
			var iter int
			if len(orders)>3 {
				iter :=	4			
			}else{
				iter := len(orders)
			}
			for j:= 0; j<=4; j++ {
				if iter >= iter{
					orderedRequests[i][j] = orders[j]
				}
				else{
					orderedRequests[i][j] = -1
				}
  			}
		}
		for k := 0; k <len(lifts);k++ {
			lifts[k].targetFloor = calculateTargetFloor()
			switch k {
				case 1:
					if lifts[k].targetFloor == -1 {
						lifts[k].targetFloor = orderedRequests[k]
					}
				case 2:
					if (lifts[k].targetFloor == -1) && () {
						
					}
			}
		}
		
	}
}*/






		/*sortedIdSlice := make([]int, aliveNodes)
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
}*/