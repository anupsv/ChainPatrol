package ChainPatrol

import (
	v0 "chainpatrol.com/v0"
	"fmt"
	"log"
)

func main() {
	configLoader := &v0.ViperConfigLoader{}
	config, err := configLoader.LoadConfig()
	if err != nil {
		log.Fatalf("Error loading config: %v", err)
	}

	if err = configLoader.CheckConfig(config); err != nil {
		log.Fatalf("Configuration check failed: %v", err)
	}

	logger, err := v0.InitializeLogger(config.LogLevel)
	if err != nil {
		log.Fatalf("Can't initialize logger: %v", err)
	}
	defer logger.Sync()

	logger.Infof("RPC Endpoint: %s", config.RPCEndpoint)
	logger.Infof("Contract Addresses: %v", config.ContractAddresses)
	logger.Infof("Log Level: %s", config.LogLevel)

	fmt.Println("ChainPatrol is running...")
}
