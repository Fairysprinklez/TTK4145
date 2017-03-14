package cost

//what kind of interface does this function have?
//it gets a map from NodeMapCompiler, but it also should update the ThisLift struct from main so that the
//FSM can do its magic. Or do you want something to bridge that?
//It could send a NodeMap both ways and main will extract the ThisLift struct? - Martin
func calculateCost()
