package portforwarderconfig

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type Config struct {
	ExternalPort        int    `json:"externalport"`
	InternalPort        int    `json:"internalport"`
	InternalIp          string `json:"internalip"`
	PortForwardDuration int    `json:"portforwardduration"`
	UpdateInterval      int    `json:"updateinterval"`
	Protocol            string `json:"protocol"`
	Description         string `json:"description"`
}

func LoadConfiguration(file string) (Config, error) {
	var config Config
	configFile, err := os.Open(file)
	if err != nil {
		return config, err
	}
	defer configFile.Close()
	jsonParser := json.NewDecoder(configFile)
	err = jsonParser.Decode(&config)
	if err != nil {
		return config, err
	}
	return config, nil
}

func ValidateConfiguration(config Config) (Config, error) {
	if config.ExternalPort == 0 {
		return config, errors.New("external port must be specified")
	}

	if config.InternalPort == 0 {
		return config, errors.New("internal port must be specified")
	}

	if len(strings.TrimSpace(config.Protocol)) == 0 {
		fmt.Println("Protocol not specified. Using TCP")
		config.Protocol = "TCP"
	}

	if len(strings.TrimSpace(config.InternalIp)) == 0 {
		ipaddress := GetOutboundIP()
		fmt.Printf("Internal IP not specified, using '%s'\n", ipaddress)
		config.InternalIp = ipaddress
	}

	if config.UpdateInterval == 0 {
		fmt.Println("Update interval not set, using 300 seconds")
		config.UpdateInterval = 300
	}

	if config.PortForwardDuration == 0 {
		fmt.Println("Port forward duration not set, using 3600 seconds")
		config.PortForwardDuration = 3600
	}

	if len(strings.TrimSpace(config.Description)) == 0 {
		desc := "Port forwarded by portforwarder"
		fmt.Printf("Description not set, using '%s'\n", desc)
		config.Description = desc
	}

	return config, nil
}

// Get preferred outbound ip of this machine
func GetOutboundIP() string {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	// TODO: why is string() vs sprintf different for IPs?
	return fmt.Sprintf("%s", localAddr.IP)
}
