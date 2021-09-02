package machine

import (
	"io/ioutil"
	"net/http"
)

func GetIP() (string, error) {
	url := "https://api.ipify.org?format=text"

	response, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	IP, ipErr := ioutil.ReadAll(response.Body)

	if ipErr != nil {
		return "", ipErr
	}

	return string(IP), nil
}
