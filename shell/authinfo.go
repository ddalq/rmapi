package shell

import (
	"github.com/abiosoft/ishell"
	"github.com/juruen/rmapi/config"
)

func authinfoCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name: "authinfo",
		Help: "display authentication tokens, user and device information",
		Func: func(c *ishell.Context) {
			configPath, err := config.ConfigPath()
			if err != nil {
				c.Printf("Error getting config path: %v\n", err)
				return
			}

			tokens := config.LoadTokens(configPath)

			c.Printf("Authentication Information:\n")
			c.Printf("========================\n")
			c.Printf("User: %s\n", ctx.UserInfo.User)
			c.Printf("Sync Version: %s\n", ctx.UserInfo.SyncVersion)
			c.Printf("\nTokens:\n")
			c.Printf("-------\n")
			
			if tokens.DeviceToken != "" {
				c.Printf("Device Token: %s...\n", truncateToken(tokens.DeviceToken))
			} else {
				c.Printf("Device Token: <not set>\n")
			}
			
			if tokens.UserToken != "" {
				c.Printf("User Token: %s...\n", truncateToken(tokens.UserToken))
			} else {
				c.Printf("User Token: <not set>\n")
			}
			
			c.Printf("\nConfig File: %s\n", configPath)
		},
	}
}

func truncateToken(token string) string {
	if len(token) > 20 {
		return token[:20]
	}
	return token
}