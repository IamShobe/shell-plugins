package npm

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func npmCLI() schema.Executable {
	return schema.Executable{
		Name:      "npm CLI", // TODO: Check if this is correct
		Runs:      []string{"npm"},
		DocsURL:   sdk.URL("https://npm.com/docs/cli"), // TODO: Replace with actual URL
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.PersonalAccessToken,
			},
		},
	}
}
