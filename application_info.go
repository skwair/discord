package harmony

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/skwair/harmony/internal/endpoint"
)

type Team struct {
	Icon        string        `json:"icon,omitempty"`
	ID          string        `json:"id,omitempty"`
	Members     []*TeamMember `json:"members,omitempty"`
	OwnerUserID string        `json:"owner_member_id,omitempty"`
}

type TeamMember struct {
	MembershipState int      `json:"membership_state,omitempty"`
	Permissions     []string `json:"permissions,omitempty"`
	TeadID          string   `json:"team_id,omitempty"`
	User            *User    `json:"user,omitempty"`
}

type ApplicationInfo struct {
	ID                  string   `json:"id,omitempty"`
	Name                string   `json:"name,omitempty"`
	Icon                string   `json:"icon,omitempty"`
	Description         string   `json:"description,omitempty"`
	RPCOrigins          []string `json:"rpc_origins,omitempty"`
	BotPublic           bool     `json:"bot_public,omitempty"`
	BotRequireCodeGrant bool     `json:"bot_require_code_grant,omitempty"`
	Owner               *User    `json:"owner,omitempty"`
	Summary             string   `json:"summary,omitempty"`
	VerifyKey           string   `json:"verify_key,omitempty"`
	Team                *Team    `json:"team,omitempty"`
	GuildID             string   `json:"guild_id,omitempty"`
	PrimarySKUID        string   `json:"primary_sku_id,omitempty"`
	Slug                string   `json:"slug,omitempty"`
	CoverImage          string   `json:"cover_image,omitempty"`
}

// ApplicationInfo returns the bot's OAuth2 application info.
func (c *Client) ApplicationInfo(ctx context.Context) (*ApplicationInfo, error) {
	e := endpoint.GetApplicationInfo()
	resp, err := c.doReq(ctx, e, nil)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, apiError(resp)
	}

	var a ApplicationInfo
	if err = json.NewDecoder(resp.Body).Decode(&a); err != nil {
		return nil, err
	}
	return &a, nil
}
