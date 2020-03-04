package cmd

import (
	"github.com/bialas1993/kubedeploy/pkg/command/converter"
	"github.com/bialas1993/kubedeploy/pkg/command/version"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var ConvertCommand = &cobra.Command{
	Use:   "convert",
	Short: "Convert template files in directory",
	Long:  `Convert template files in directory`,
	Run: func(cmd *cobra.Command, args []string) {
		out, _ := cmd.Flags().GetString("output")
		tpl, _ := cmd.Flags().GetString("templates")
		env, _ := cmd.Flags().GetString("env")

		if err := converter.New(out, tpl, env, version.Hash()).Do(); err != nil {
			log.Error(err)
			return
		}

		log.Infof("Converting completed.")
	},
}
