package alfacrm

const (
	httpURL    = "https://"
	alfacrmURL = "www.alfacrm.ru"
	xToken     = "X-ALFACRM-TOKEN"
)

const (
	authURL = "v2api/auth/login"
)

type makeRequestOptions struct {
	// Method is a request's method
	Method string
	// BaseURL is a url from constants above.
	BaseURL string
	// In is a struct, which will be marshalled to Request Body
	In interface{}
	// Out is a struct, which will be unmarshalled
	Out interface{}
	// Params is a URL Parameters
	Params *Params
}

type Params struct {
	With      With
	Page      string
	Limit     string
	Query     string
	Filter    string
	Order     string
	ContactID string
	ChatID    string
}

type With struct {
	CatalogElements bool
	Contacts        bool
	Leads           bool
	Companies       bool
}
