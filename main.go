package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fightingsleep/portforwarder/portforwarder"
	"github.com/fightingsleep/portforwarder/portforwarderconfig"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide config filename argument")
	}

	// Read the config
	config, err := portforwarderconfig.LoadConfiguration(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	config, err = portforwarderconfig.ValidateConfiguration(config)
	if err != nil {
		log.Fatal(err)
	}

	for {
		// Forward the port every x seconds
		err = portforwarder.GetIPAndForwardPort(context.Background(), config)
		if err != nil {
			log.Fatal(err)
		}
		time.Sleep(time.Second * time.Duration(config.UpdateInterval))
	}
}
