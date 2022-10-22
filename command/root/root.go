package root

import (
	"fmt"
	"os"

	"github.com/alveycoin/alveychain/command/backup"
	"github.com/alveycoin/alveychain/command/genesis"
	"github.com/alveycoin/alveychain/command/helper"
	"github.com/alveycoin/alveychain/command/ibft"
	"github.com/alveycoin/alveychain/command/license"
	"github.com/alveycoin/alveychain/command/loadbot"
	"github.com/alveycoin/alveychain/command/monitor"
	"github.com/alveycoin/alveychain/command/peers"
	"github.com/alveycoin/alveychain/command/secrets"
	"github.com/alveycoin/alveychain/command/server"
	"github.com/alveycoin/alveychain/command/status"
	"github.com/alveycoin/alveychain/command/txpool"
	"github.com/alveycoin/alveychain/command/version"
	"github.com/alveycoin/alveychain/command/whitelist"
	"github.com/spf13/cobra"
)

type RootCommand struct {
	baseCmd *cobra.Command
}

func NewRootCommand() *RootCommand {
	rootCommand := &RootCommand{
		baseCmd: &cobra.Command{
			Short: "Alvey Coin is a framework for building Ethereum-compatible Blockchain networks",
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
		loadbot.GetCommand(),
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
