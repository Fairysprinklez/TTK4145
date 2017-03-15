package targetFloorAssigner

import (
	"../config"
	"math"
	"sort"
)

type lifts []config.Lift



func (slice lifts) Len() int{
	return len(slice)
}

func (slice lifts) Less(i, j int) bool{
	return slice[i].ID < slice[j].ID
}

func (slice lifts) Swap(i, j int){
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
	for f = config.NumFloors-1; f >= lift.LastKnownFloor; f-- {
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
		requests = append(requests, requestsAbove...)
	case config.MD_Stop:
		if len(requestsAbove)>0 && len(requestsBelow)>0 {
			distanceUp := math.Abs(float64(lift.LastKnownFloor - requestsAbove[0]))
			distanceDown := math.Abs(float64(lift.LastKnownFloor - requestsBelow[0]))
			if distanceUp < distanceDown{
				requests = requestsAbove
				requests = append(requests, requestsBelow...)
			}else {
				requests = requestsBelow
				requests = append(requests, requestsAbove...)
			}
		}else{
			requests = requestsBelow
			requests = append(requests, requestsAbove...)
		}
		
		
	}
	return requests
}


func TargetFloorAssigner(nodeMapInCh chan config.NodeMap, thisLiftOutCh chan config.Lift, doneReadingMap chan bool, myID string) {
	var lifts lifts
	var validHallRequests [config.NumFloors][config.NumButtons-1]bool
	for {
		for f := 0; f < config.NumFloors; f++ {
			for b := 0; b < config.NumButtons-1; b++ {
				validHallRequests[f][b] = true 
			}
		}

		nodeMap := <- nodeMapInCh
		aliveNodes := 0
		for id := range nodeMap {
			liftNode := nodeMap[id]
			if liftNode.Alive {
				lifts = append(lifts, liftNode)
				aliveNodes++
				for f := 0; f < config.NumFloors; f++ {
					for b := 0; b < config.NumButtons-1; b++ {
						validHallRequests[f][b] = (validHallRequests[f][b] && liftNode.Requests[f][b])
					}
				}
			}
		}
		lightMatrix := validHallRequests
		doneReadingMap <- true
		sort.Sort(lifts)
		var target int
		for i := 0; i < len(lifts); i++ {
			requests := prioritiseOrders(lifts[i], validHallRequests)
			if len(requests)>0{
				target = requests[1]
			}else{
				target = -1
			}
			lifts[i].TargetFloor = target
			if lifts[i].ID == myID{
				for f := 0; f < config.NumFloors; f++{
					lifts[i].Requests[f][config.B_HallUp] = lightMatrix[f][config.B_HallUp]
					lifts[i].Requests[f][config.B_HallDown] = lightMatrix[f][config.B_HallDown]
				}
				thisLiftOutCh <- lifts[i]
			}
			if target != -1{
				validHallRequests[target][config.B_HallUp] = false
				validHallRequests[target][config.B_HallDown] = false
			}
		}
		lifts = nil
	}

}
