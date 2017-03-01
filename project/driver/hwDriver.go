package driver


const NumFloors int = 4
const NumButtons int = 3
const MotorSpeed int = 2800

var lampChannelMatrix = [NumFloors][NumButtons]int{
	{LIGHT_UP1, LIGHT_DOWN1, LIGHT_COMMAND1},
	{LIGHT_UP2, LIGHT_DOWN2, LIGHT_COMMAND2},
	{LIGHT_UP3, LIGHT_DOWN3, LIGHT_COMMAND3},
	{LIGHT_UP4, LIGHT_DOWN4, LIGHT_COMMAND4},
}

var buttonChannelMatrix = [NumFloors][NumButtons]int{
	{BUTTON_UP1, BUTTON_DOWN1, BUTTON_COMMAND1},
	{BUTTON_UP2, BUTTON_DOWN2, BUTTON_COMMAND2},
	{BUTTON_UP3, BUTTON_DOWN3, BUTTON_COMMAND3},
	{BUTTON_UP4, BUTTON_DOWN4, BUTTON_COMMAND4},
}

func ElevInit() int {
    if !io_init() {
    	return -1
    }

    for f := 0; f < NumFloors; f++ {
    	for b := 0; b <NumButtons; b++ {
    		ElevSetButtonLamp(f, b, false)
    	}
    }

    
    ElevSetStopLamp(false);
    ElevSetDoorOpenLamp(false);
    ElevSetFloorIndicator(0);

    return 0
}


func ElevSetMotorDirection(dirn int) {
    if dirn == 0{
        io_write_analog(MOTOR, 0)
    } else if dirn > 0 {
        io_clear_bit(MOTORDIR)
        io_write_analog(MOTOR, MotorSpeed)
    } else if dirn < 0 {
        io_set_bit(MOTORDIR)
        io_write_analog(MOTOR, MotorSpeed)
    }
}


func ElevSetButtonLamp(floor int, button int, value bool) {
    //TODO: add functionality to check valid input

    if value {
        io_set_bit(lampChannelMatrix[floor][button])
    } else {
        io_clear_bit(lampChannelMatrix[floor][button])
    }
}


func ElevSetFloorIndicator(floor int) {
    //TODO: add functionality to check valid input


    // Binary encoding. One light must always be on.
    if floor&0x02 > 0 {
        io_set_bit(LIGHT_FLOOR_IND1)
    } else {
        io_clear_bit(LIGHT_FLOOR_IND1)
    }    

    if floor&0x01 > 0 {
        io_set_bit(LIGHT_FLOOR_IND2)
    } else {
        io_clear_bit(LIGHT_FLOOR_IND2)
    }    
}


func ElevSetDoorOpenLamp(value bool) {
    if value {
        io_set_bit(LIGHT_DOOR_OPEN)
    } else {
        io_clear_bit(LIGHT_DOOR_OPEN)
    }
}


func ElevSetStopLamp(value bool) {
    if value {
        io_set_bit(LIGHT_STOP)
    } else {
        io_clear_bit(LIGHT_STOP)
    }
}



func ElevGetButtonSignal(floor int, button int) bool {
    //TODO: add functionality to check valid input

	if io_read_bit(buttonChannelMatrix[floor][button]) {
		return true
	} else {
		return false
	}
}


func ElevGetFloorSensorSignal() int {
    if io_read_bit(SENSOR_FLOOR1) {
        return 0
    } else if io_read_bit(SENSOR_FLOOR2) {
        return 1
    } else if io_read_bit(SENSOR_FLOOR3) {
        return 2
    } else if io_read_bit(SENSOR_FLOOR4) {
        return 3
    } else {
        return -1
    }
}


func ElevGetStopSignal() bool {
    return io_read_bit(STOP)
}


func ElevGetObstructionSignal() bool {
    return io_read_bit(OBSTRUCTION)
}