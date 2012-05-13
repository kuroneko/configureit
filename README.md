# ConfigureIt

	import (
		"github.com/kuroneko/configureit"
		"os"
		"fmt"
	)
	
	var config *configureit.Config
	
	func DoConfiguration(configfile string) {
		config = configureit.New()
		config.Add("key_identifier", configureit.NewStringOption("Default value"))
		fh, _ := os.Open(configfile)
		config.Read(fh, 1)
	}
	
	func UseConfig() {
		cn := config.Get("key_identifier")
		// do stuff with the confignode...
		fmt.Printf("key_identifier = %s\n", cn)
	}

ConfigureIt implements a simple line-oriented configuration file
parser.

The ConfigureIt parser is modelled after a design I have personally
used time and time again in C projects, appropriately genercised so
allow arbitrary types to be easily extended into the system.

It is important to note that the API is not stable at this time and
may change at a whim in order to make it more idiomatic.


