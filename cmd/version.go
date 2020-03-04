package cmd

import (
	"fmt"
	"os"

	"github.com/bialas1993/kubedeploy/pkg/command/version"
	"github.com/spf13/cobra"
)

var VersionCommand = &cobra.Command{
	Use:   "version",
	Short: "Hash version",
	Long:  `Hash version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Fprintf(os.Stdout, version.Hash())
	},
}
