package root

import (
	"fmt"
	"os"

	"github.com/ksandr84on/nera/command/backup"
	"github.com/ksandr84on/nera/command/genesis"
	"github.com/ksandr84on/nera/command/helper"
	"github.com/ksandr84on/nera/command/ibft"
	"github.com/ksandr84on/nera/command/license"
	"github.com/ksandr84on/nera/command/monitor"
	"github.com/ksandr84on/nera/command/peers"
	"github.com/ksandr84on/nera/command/secrets"
	"github.com/ksandr84on/nera/command/server"
	"github.com/ksandr84on/nera/command/status"
	"github.com/ksandr84on/nera/command/txpool"
	"github.com/ksandr84on/nera/command/version"
	"github.com/ksandr84on/nera/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Greenera network. Every transaction helps the planet",
		},
	}

	helper.RegisterJSONOutputFlag(rootCommand.baseCmd)

	rootCommand.registerSubCommands()

	return rootCommand
}

func (rc *RootCommand) registerSubCommands() {
	rc.baseCmd.AddCommand(
		version.GetCommand(),
		txpool.GetCommand(),
		status.GetCommand(),
		secrets.GetCommand(),
		peers.GetCommand(),
		monitor.GetCommand(),
		ibft.GetCommand(),
		backup.GetCommand(),
		genesis.GetCommand(),
		server.GetCommand(),
		whitelist.GetCommand(),
		license.GetCommand(),
	)
}

func (rc *RootCommand) Execute() {
	if err := rc.baseCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)

		os.Exit(1)
	}
}
