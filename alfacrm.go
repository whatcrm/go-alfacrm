package alfacrm

import (
	"fmt"
)

func NewAPI(apiKey string) *API {
	return &API{
		ApiKey: apiKey,
		Debug:  false,
	}
}

func (api *API) SetOptions(domain, apiKey string, debug bool) error {
	if isRegex(domain) {
		api.Domain = domain
	} else {
		return fmt.Errorf("domain name is not set")
	}

	if apiKey != "" {
		api.ApiKey = apiKey
	} else {
		return fmt.Errorf("API Key is not set")
	}

	api.Debug = debug

	return nil
}
