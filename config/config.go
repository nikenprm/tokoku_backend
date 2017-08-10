package config

import (
	"encoding/json"
	"os"
)

//Configuration is a place for placing the config in a json format
type Configuration struct {
	DB      DBParam      `json:"db"`
	Walmart WalmartParam `json:"walmart"`
	Email   EmailParam   `json:"email"`
}

//Config variables for initiate new Configuration
var Config Configuration

//LoadConfig is used for load configuration
func LoadConfig(configPath string) (err error) {
	var file *os.File
	if file, err = os.Open(configPath); err != nil {
		return err
	}
	if err = json.NewDecoder(file).Decode(&Config); err != nil {
		return err
	}
	return nil
}
