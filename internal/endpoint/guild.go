package endpoint

func CreateGuild() *Endpoint {
	return &Endpoint{
		URL: "/guilds",
		Key: "/guilds",
	}
}

func GetGuild(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID,
		Key: "/guilds/" + guildID,
	}
}

func ModifyGuild(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID,
		Key: "/guilds/" + guildID,
	}
}

func DeleteGuild(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID,
		Key: "/guilds/" + guildID,
	}
}

func GetGuildChannels(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/channels",
		Key: "/guilds/" + guildID + "/channels",
	}
}

func CreateGuildChannel(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/channels",
		Key: "/guilds/" + guildID + "/channels",
	}
}

func ModifyChannelPositions(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/channels",
		Key: "/guilds/" + guildID + "/channels",
	}
}

func GetGuildMember(guildID, userID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members/" + userID,
		Key: "/guilds/" + guildID + "/members",
	}
}

func GetGuildMembers(guildID, query string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members?" + query,
		Key: "/guilds/" + guildID + "/members",
	}
}

func AddGuildMember(guildID, userID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members/" + userID,
		Key: "/guilds/" + guildID + "/members",
	}
}

func RemoveGuildMember(guildID, userID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members/" + userID,
		Key: "/guilds/" + guildID + "/members",
	}
}

func ModifyGuildMember(guildID, userID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members/" + userID,
		Key: "/guilds/" + guildID + "/members",
	}
}

func ModifyCurrentUserNick(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/members/@me/nick",
		Key: "/guilds/" + guildID + "/members/@me/nick",
	}
}

func GetGuildBans(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/bans",
		Key: "/guilds/" + guildID + "/bans",
	}
}

func BanWithReason(guildID, userID, query string) *Endpoint {
	if query != "" {
		query = "?" + query
	}

	return &Endpoint{
		URL: "/guilds/" + guildID + "/bans/" + userID + query,
		Key: "/guilds/" + guildID + "/bans",
	}
}

func Unban(guildID, userID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/bans/" + userID,
		Key: "/guilds/" + guildID + "/bans",
	}
}

func GetPruneCount(guildID, query string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/prune?" + query,
		Key: "/guilds/" + guildID + "/prune",
	}
}

func BeginPrune(guildID, query string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/prune?" + query,
		Key: "/guilds/" + guildID + "/prune",
	}
}

func GetGuildVoiceRegions(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/regions",
		Key: "/guilds/" + guildID + "/regions",
	}
}

func GetGuildInvites(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/invites",
		Key: "/guilds/" + guildID + "/invites",
	}
}

func GetGuildIntegrations(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/integrations",
		Key: "/guilds/" + guildID + "/integrations",
	}
}

func AddGuildIntegration(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/integrations",
		Key: "/guilds/" + guildID + "/integrations",
	}
}

func ModifyGuildIntegration(guildID, integrationID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/integrations/" + integrationID,
		Key: "/guilds/" + guildID + "/integrations",
	}
}

func RemoveGuildIntegration(guildID, integrationID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/integrations/" + integrationID,
		Key: "/guilds/" + guildID + "/integrations",
	}
}

func SyncGuildIntegration(guildID, integrationID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/integrations/" + integrationID + "/sync",
		Key: "/guilds/" + guildID + "/integrations",
	}
}

func GetGuildEmbed(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/embed",
		Key: "/guilds/" + guildID + "/embed",
	}
}

func ModifyGuildEmbed(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/embed",
		Key: "/guilds/" + guildID + "/embed",
	}
}

func GetGuildVanityURL(guildID string) *Endpoint {
	return &Endpoint{
		URL: "/guilds/" + guildID + "/vanity-url",
		Key: "/guilds/" + guildID + "/vanity-url",
	}
}
