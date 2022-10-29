package portician

import (
	"encoding/json"
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

type PortForwardConfig struct {
	ExternalPort        int    `json:"externalport"`
	InternalPort        int    `json:"internalport"`
	InternalIp          string `json:"internalip"`
	PortForwardDuration int    `json:"portforwardduration"`
	Protocol            string `json:"protocol"`
	Description         string `json:"description"`
}

type Config struct {
	UpdateInterval int                 `json:"updateinterval"`
	Configs        []PortForwardConfig `json:"configs"`
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

func ValidateConfiguration(config *Config) error {
	if config.UpdateInterval == 0 {
		fmt.Println("🤌  Update interval not set, using 300 seconds")
		config.UpdateInterval = 300
	}

	for index, _ := range config.Configs {
		if config.Configs[index].ExternalPort == 0 {
			return fmt.Errorf("🚨  error in config '%d'. external port must be specified  🚨", index)
		}

		if config.Configs[index].InternalPort == 0 {
			return fmt.Errorf("🚨  error in config '%d'. internal port must be specified  🚨", index)
		}

		if len(strings.TrimSpace(config.Configs[index].Protocol)) == 0 {
			fmt.Printf("🤌  Protocol not specified in config '%d'. Using TCP\n", index)
			config.Configs[index].Protocol = "TCP"
		}

		if len(strings.TrimSpace(config.Configs[index].InternalIp)) == 0 {
			ipaddress := GetOutboundIP()
			fmt.Printf("🤌  Internal IP not specified in config '%d', using '%s'\n", index, ipaddress)
			config.Configs[index].InternalIp = ipaddress
		}

		if config.Configs[index].PortForwardDuration == 0 {
			fmt.Printf("🤌  Port forward duration not set in config '%d', using 3600 seconds\n", index)
			config.Configs[index].PortForwardDuration = 3600
		}

		if len(strings.TrimSpace(config.Configs[index].Description)) == 0 {
			desc := "Port forwarded by portician"
			fmt.Printf("🤌  Description not set in config '%d', using '%s'\n", index, desc)
			config.Configs[index].Description = desc
		}
	}

	return nil
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
