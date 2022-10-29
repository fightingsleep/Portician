package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/fightingsleep/portician/pkg/portician"
)

func main() {
	var configFilePath string = ""
	if len(os.Args) != 2 {
		fmt.Println("🤌  Config file path not specified. Assuming config.json")
		configFilePath = "config.json"
	} else {
		configFilePath = os.Args[1]
	}

	// Read the config
	configuration, err := portician.LoadConfiguration(configFilePath)
	if err != nil {
		log.Fatal(err)
	}

	err = portician.ValidateConfiguration(&configuration)
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
			fmt.Printf(
				"🍔  Port '%d' forwarded to '%s:%d' successfully  🍔\n",
				config.ExternalPort,
				config.InternalIp,
				config.InternalPort)
		}
		time.Sleep(time.Second * time.Duration(configuration.UpdateInterval))
	}
}
