//TODO: package

import "./config"
import "./network"


/*TODO: determine channels:
	
	
*/
func universeCompiler(/*channel input*/){
	//Initializing the map of the universe
	universe = make(map[string]config.Lift)
	universe[network.id] = Lift{
		network.id,
		-1, 
		0, 
		{{false, false},
		 {false, false},
		 {false, false},
		 {false, false}},
		{false, false, false, false}
	}
	
	
	