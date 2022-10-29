package main

import (
	"context"
	"log"
	"os"
	"time"

	"github.com/fightingsleep/portician/portician"
	conf "github.com/fightingsleep/portician/porticianconfig"
)

func main() {
	if len(os.Args) != 2 {
		log.Fatal("Must provide config filename argument")
	}

	// Read the config
	configuration, err := conf.LoadConfiguration(os.Args[1])
	if err != nil {
		log.Fatal(err)
	}

	err = conf.ValidateConfiguration(&configuration)
	if err != nil {
		log.Fatal(err)
	}

	for {
		for _, config := range configuration.Configs {
			// Forward the port every x seconds
			err = portician.ForwardPort(context.Background(), config)
			if err != nil {
				log.Fatal(err)
			}
		}
		time.Sleep(time.Second * time.Duration(configuration.UpdateInterval))
	}
}
