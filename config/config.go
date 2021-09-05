package config

import (
	"fmt"
	"time"
)

type Config struct {
	Id            string        `json:"id"`
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

type VersionUpdate struct {
	Update bool   `json:"update"`
	Url    string `json:"url"`
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
