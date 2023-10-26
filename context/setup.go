package context

import (
	"GoCrudORM/types"
	"encoding/json"
	"flag"
	"os"
	"strings"
)

func LoadConfig() (*types.Configuration, error) {
	var environment string
	flag.StringVar(&environment, "env", "development", "Environment (development/production)")
	flag.Parse()

	file, err := os.Open("config.json")
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var config map[string]types.Configuration
	if err := json.NewDecoder(file).Decode(&config); err != nil {
		return nil, err
	}

	selectedConfig := config[strings.ToLower(environment)]
	return &selectedConfig, nil
}
