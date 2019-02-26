package main

import (
	"context"
	"fmt"
	"os"

	"github.com/skwair/harmony"
	"github.com/skwair/harmony/audit"
)

func main() {
	token := os.Getenv("BOT_TOKEN")
	if token == "" {
		fmt.Fprint(os.Stderr, "Environment variable BOT_TOKEN must be set.")
		return
	}

	// The guild ID you want to get the audit log from.
	// Requires the bot to have the 'VIEW_AUDIT_LOG' permission.
	guildID := os.Getenv("GUILD_ID")
	if guildID == "" {
		fmt.Fprint(os.Stderr, "Environment variable GUILD_ID must be set.")
	}

	client, err := harmony.NewClient(token)
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	log, err := client.AuditLog(context.Background(), guildID, harmony.WithLimit(25))
	if err != nil {
		fmt.Fprint(os.Stderr, err)
		return
	}

	printRoleEntries(log)
}

func printRoleEntries(log *audit.Log) {
	// The audit log is composed of entries.
	for _, entry := range log.Entries {
		// Each entry has a type, matching the type of event this entry describes.
		switch e := entry.(type) {
		case *audit.RoleCreate:
			fmt.Printf("role %q was created", e.Name)
			// Here, the entry will contain all the settings this role was created with.

		case *audit.RoleUpdate:
			fmt.Printf("role %q was updated", e.Name)

			// Fields that are of type *StringValues, *IntValues, *BoolValues
			// are settings that have potentially been modified. If they are non-nil,
			// it means they were and they will hold the old as well as the new value.
			if e.Name != nil {
				fmt.Printf("name changed from %q to %q", e.Name.Old, e.Name.New)
			}

		case *audit.RoleDelete:
			fmt.Printf("role %q was deleted", e.Name)
			// Here, the entry will contain all the settings this role had before being deleted.
		}
	}
}
