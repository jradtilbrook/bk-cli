package root

import (
	"github.com/MakeNowJust/heredoc"
	authCmd "github.com/buildkite/cli/v3/pkg/cmd/auth"
	versionCmd "github.com/buildkite/cli/v3/pkg/cmd/version"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func NewCmdRoot(viper *viper.Viper, version string) (*cobra.Command, error) {
	cmd := &cobra.Command{
		Use:   "bk <command> <subcommand> [flags]",
		Short: "Buildkite CLI",
		Long:  "Work with Buildkite from the command line.",
		Example: heredoc.Doc(`
			$ bk build view
		`),
		Annotations: map[string]string{
			"versionInfo": versionCmd.Format(version),
		},
	}

	cmd.AddCommand(versionCmd.NewCmdVersion())
	cmd.AddCommand(authCmd.NewCmdAuth(viper))

	return cmd, nil
}