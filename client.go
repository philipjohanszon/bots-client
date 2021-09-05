package main

import (
	"sync"
	"time"

	_config "github.com/philipjohanszon/bots-client/config"
	_network "github.com/philipjohanszon/bots-client/network"
)

var (
	networkConfig = _config.NetworkConfig{
		MasterURL: "http://localhost:3000",
		Timer:     time.Second * 2,
	}

	config = _config.Config{
		Version:       "0.01",
		NetworkConfig: networkConfig,
	}

	network = _network.Network{
		Config: config,
	}
)

func main() {
	var wg sync.WaitGroup
	wg.Add(1)
	go network.Start()
	wg.Wait()
}
