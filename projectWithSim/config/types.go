package config

type ElevatorType int
const (
    ET_Comedi ElevatorType  = 0
    ET_Simulation           = 1
)

type MotorDirection int
const (
    MD_Up   MotorDirection  = 1
    MD_Down                 = -1
    MD_Stop                 = 0
)

type ButtonType int
const (
    B_HallUp = iota
    B_HallDown
    B_Cab
)
    

type ButtonEvent struct {
    Floor  int
    Button ButtonType
}
