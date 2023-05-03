package alfacrm

// Handler with main data to work with amoCRM

type API struct {
	ApiKey      string
	Domain      string
	RedirectURI string
	Debug       bool
}

type RequestAnswer struct {
	Name    string `json:"name"`
	Message string `json:"message"`
	Code    int    `json:"code"`
	Status  int    `json:"status"`
}

type TokenReciever struct {
	Token string `json:"token"`
}
