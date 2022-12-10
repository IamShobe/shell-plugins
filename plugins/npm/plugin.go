package npm

import (
	"github.com/1Password/shell-plugins/sdk"
	"github.com/1Password/shell-plugins/sdk/schema"
)

func New() schema.Plugin {
	return schema.Plugin{
		Name: "npm",
		Platform: schema.PlatformInfo{
			Name:     "npm",
			Homepage: sdk.URL("https://npm.com"), // TODO: Check if this is correct
		},
		Credentials: []schema.CredentialType{
			PersonalAccessToken(),
		},
		Executables: []schema.Executable{
			npmCLI(),
		},
	}
}
