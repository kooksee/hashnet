package cmds

import (
	"fmt"
	"github.com/kooksee/hashnet/version"
	"github.com/spf13/cobra"
	tv "github.com/tendermint/tendermint/version"
)

// VersionCmd ...
var VersionCmd = &cobra.Command{
	Use:     "version",
	Aliases: []string{"v", "ver"},
	Short:   "Show Version Info",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("tendermint version", tv.Version)
		fmt.Println("hashnet version", version.Version)
		fmt.Println("hashnet commit version", version.GitCommit)
		fmt.Println("hashnet build version", version.BuildVersion)
	},
}
