package main

import (
	"github.com/kooksee/hashnet/app"
	"github.com/kooksee/hashnet/cmd"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/cli"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
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
	rootCmd.AddCommand(commands.NewRunNodeCmd(func(config *config.Config, logger log.Logger) (i *node.Node, e error) {
		// Generate node PrivKey
		nodeKey, err := p2p.LoadOrGenNodeKey(config.NodeKeyFile())
		if err != nil {
			return nil, err
		}
		return node.NewNode(config,
			privval.LoadOrGenFilePV(config.PrivValidatorFile()),
			nodeKey,
			proxy.NewLocalClientCreator(app.New(logger)),
			node.DefaultGenesisDocProviderFunc(config),
			node.DefaultDBProvider,
			node.DefaultMetricsProvider(config.Instrumentation),
			logger,
		)
	}))
	if err := cli.PrepareBaseCmd(rootCmd, "HASHNET", os.ExpandEnv("$PWD/kdata")).Execute(); err != nil {
		panic(err)
	}
}
