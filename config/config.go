package config

import (
	"fmt"
	"time"
)

type Config struct {
	Version       string        `json:"version"`
	NetworkConfig NetworkConfig `json:"networkConfig"`
}

type ConfigUpdate struct {
	Type  string `json:"type"`
	Field string `json:"field"`
	Value string `json:"value"`
}

type NetworkConfig struct {
	Timer     time.Duration `json:"networkTimer"`
	MasterURL string        `json:"masterUrl"`
}

func (config *Config) SaveConfig() error {

	return nil
}

func (config *Config) ConfigUpdate(channel chan ConfigUpdate) {

	for update := range channel {
		switch update.Type {
		case "networkConfig":
			fmt.Println("updated network config")
			fmt.Println(update)
		}
	}
}
