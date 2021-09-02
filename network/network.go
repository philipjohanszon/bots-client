package network

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/philipjohanszon/bots-client/config"
	"github.com/philipjohanszon/bots-client/machine"
)

type Network struct {
	Config config.Config
}

//loops through with the speed set by the duration and sends healthcheck to the server and also gets command/config update data
func (network *Network) Start() {
	for true {
		go network.getUpdate()
		time.Sleep(network.Config.NetworkConfig.Timer)
	}
}

func (network *Network) getUpdate() {

	jsonData, dataError := createJSONData(&network.Config)

	if dataError != nil {
		fmt.Println(dataError.Error())
		return
	}

	response, responseError := http.Post(network.Config.NetworkConfig.MasterURL+"/bots/update", "application/json", bytes.NewBuffer(*jsonData))

	if responseError != nil {
		fmt.Println(responseError.Error())
		return
	}

	defer response.Body.Close()

}

/*Creating the data that will be sent to the server

  Docs say that it should contain (sample values):

  {
      "version": "1.05",
      "ip": "94.36.356.23",
      "macAdress": "f3:ce:76:d2:49:0d",
      "config": {
          "masterUrl"
          "networkTimer": 1200000,
      },
  }
*/

type JSONData struct {
	Version string `json:"version"`
	Ip      string `json:"ip"`
	// MacAdress string        `json:"macAdress"`
	Config config.Config `json:"config"`
}

func createJSONData(configStruct *config.Config) (*[]byte, error) {

	data := JSONData{}
	data.Config = *configStruct

	//Need to get version, Ip and MAC Adress
	data.Version = configStruct.Version
	IP, ipError := machine.GetIP()

	if ipError != nil {
		fmt.Println(ipError.Error())
		return nil, nil
	}

	data.Ip = IP

	jsonData, marshallErr := json.Marshal(data)

	if marshallErr != nil {
		fmt.Println(marshallErr.Error())
		return nil, nil
	}

	fmt.Println("json")
	fmt.Println(string(jsonData))

	return &jsonData, nil

}
