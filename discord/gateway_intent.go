package discord

// GatewayIntent specifies which events the Gateway should send to a client.
type GatewayIntent int

// String implements fmt.Stringer.
func (i GatewayIntent) String() string {
	return intentNames[i]
}

// List of gateway intents a client can subscribe to.
const (
	GatewayIntentGuild                  GatewayIntent = 1 << 0
	GatewayIntentGuildMembers           GatewayIntent = 1 << 1
	GatewayIntentGuildBans              GatewayIntent = 1 << 2
	GatewayIntentGuildEmojis            GatewayIntent = 1 << 3
	GatewayIntentGuildIntegrations      GatewayIntent = 1 << 4
	GatewayIntentGuildWebhooks          GatewayIntent = 1 << 5
	GatewayIntentGuildInvites           GatewayIntent = 1 << 6
	GatewayIntentGuildVoiceStates       GatewayIntent = 1 << 7
	GatewayIntentGuildPresences         GatewayIntent = 1 << 8
	GatewayIntentGuildMessages          GatewayIntent = 1 << 9
	GatewayIntentGuildMessageReactions  GatewayIntent = 1 << 10
	GatewayIntentGuildMessageTyping     GatewayIntent = 1 << 11
	GatewayIntentDirectMessages         GatewayIntent = 1 << 12
	GatewayIntentDirectMessageReactions GatewayIntent = 1 << 13
	GatewayIntentDirectMessageTyping    GatewayIntent = 1 << 14
)

// Equivalent to all intents except privileged (GatewayIntentGuildMembers and GatewayIntentGuildPresences), OR'd.
const GatewayIntentUnprivileged = GatewayIntentGuild | GatewayIntentGuildBans | GatewayIntentGuildEmojis | GatewayIntentGuildIntegrations | GatewayIntentGuildWebhooks | GatewayIntentGuildInvites | GatewayIntentGuildVoiceStates | GatewayIntentGuildMessages | GatewayIntentGuildMessageReactions | GatewayIntentGuildMessageTyping | GatewayIntentDirectMessages | GatewayIntentDirectMessageReactions | GatewayIntentDirectMessageTyping

var intentNames = map[GatewayIntent]string{
	GatewayIntentGuild:                  "GUILDS",
	GatewayIntentGuildMembers:           "GUILD_MEMBERS",
	GatewayIntentGuildBans:              "GUILD_BANS",
	GatewayIntentGuildEmojis:            "GUILD_EMOJIS",
	GatewayIntentGuildIntegrations:      "GUILD_INTEGRATIONS",
	GatewayIntentGuildWebhooks:          "GUILD_WEBHOOKS",
	GatewayIntentGuildInvites:           "GUILD_INVITES",
	GatewayIntentGuildVoiceStates:       "GUILD_VOICE_STATES",
	GatewayIntentGuildPresences:         "GUILD_PRESENCES",
	GatewayIntentGuildMessages:          "GUILD_MESSAGES",
	GatewayIntentGuildMessageReactions:  "GUILD_MESSAGE_REACTIONS",
	GatewayIntentGuildMessageTyping:     "GUILD_MESSAGE_TYPING",
	GatewayIntentDirectMessages:         "DIRECT_MESSAGES",
	GatewayIntentDirectMessageReactions: "DIRECT_MESSAGE_REACTIONS",
	GatewayIntentDirectMessageTyping:    "DIRECT_MESSAGE_TYPING",
}
