package main

import (
	"github.com/kooksee/hashnet/cmd"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/node"
	"os"

	"github.com/tendermint/tendermint/cmd/tendermint/commands"
)

func main() {
	rootCmd := commands.RootCmd
	rootCmd.AddCommand(
		commands.GenValidatorCmd,
		commands.InitFilesCmd,
		commands.ProbeUpnpCmd,
		commands.LiteCmd,
		commands.ReplayCmd,
		commands.ReplayConsoleCmd,
		commands.ResetAllCmd,
		commands.ResetPrivValidatorCmd,
		commands.ShowValidatorCmd,
		commands.TestnetFilesCmd,
		commands.ShowNodeIDCmd,
		commands.GenNodeKeyCmd,
		cmds.VersionCmd,
	)

	// Create & start node
	rootCmd.AddCommand(commands.NewRunNodeCmd(node.DefaultNewNode))
	if err := cli.PrepareBaseCmd(rootCmd, "HASHNET", os.ExpandEnv("$PWD/kdata")).Execute(); err != nil {
		panic(err)
	}
}
