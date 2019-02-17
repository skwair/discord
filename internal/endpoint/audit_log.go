package endpoint

import "net/http"

func GetAuditLog(guildID, query string) *Endpoint {
	if query != "" {
		query = "?" + query
	}

	return &Endpoint{
		Method: http.MethodGet,
		URL:    "/guilds/" + guildID + "/audit-logs" + query,
		Key:    "/guilds/" + guildID + "/audit-logs",
	}
}
