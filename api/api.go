package api

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

const base = "https://api.digitalocean.com"

var (
	clientId string
	apiKey   string
)

func init() {
	clientId = os.Getenv("DIGITAL_OCEAN_CLIENT_ID")
	apiKey = os.Getenv("DIGITAL_OCEAN_API_KEY")
	exit := false
	if clientId == "" {
		fmt.Fprintln(os.Stderr, "Please set your DIGITAL_OCEAN_CLIENT_ID environment variable")
		exit = true
	}
	if apiKey == "" {
		fmt.Fprintln(os.Stderr, "Please set your DIGITAL_OCEAN_API_KEY environment variable")
		exit = true
	}
	if exit {
		os.Exit(2)
	}
}

func GetUrl(resource, id, action string) string {
	var s string
	switch action {
	case "list":
		s = fmt.Sprintf("%v/%v", base, resource)
	case "create":
		s = fmt.Sprintf("%v/%v/new", base, resource)
	default:
		s = fmt.Sprintf("%v/%v/%v/%v", base, resource, id, action)
	}
	return AddCredentials(s)
}

func AddCredentials(s string) string {
	return fmt.Sprintf("%v?client_id=%v&api_key=%v", s, clientId, apiKey)
}

func MakeRequest(URL string) ([]byte, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	resp.Body.Close()
	return bytes, nil
}

func Call(resource, id, action string) ([]byte, error) {
	return MakeRequest(GetUrl(resource, id, action))
}
