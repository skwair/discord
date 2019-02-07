package endpoint

import "net/http"

func Gateway() *Endpoint {
	return &Endpoint{
		Method: http.MethodGet,
		URL:    "/gateway",
		Key:    "/gateway",
	}
}

func GatewayBot() *Endpoint {
	return &Endpoint{
		Method: http.MethodGet,
		URL:    "/gateway/bot",
		Key:    "/gateway/bot",
	}
}
