package npm

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/needsauth"
	"github.com/1Password/shell-plugins/sdk/schema"
	"github.com/1Password/shell-plugins/sdk/schema/credname"
)

func npmCLI() schema.Executable {
	return schema.Executable{
		Name:      "npm CLI",
		Runs:      []string{"npm"},
		DocsURL:   sdk.URL("https://docs.npmjs.com/cli/"),
		NeedsAuth: needsauth.NotForHelpOrVersion(),
		Uses: []schema.CredentialUsage{
			{
				Name: credname.PersonalAccessToken,
			},
		},
	}
}
