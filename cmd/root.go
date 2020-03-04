package cmd

import (
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "kubedeploy",
	Short: "Deploy tool for jenkins",
	Long:  `Deploy tool for jenkins integrate kubernetes API.`,
}

func init() {
	flagsBind(ConvertCommand)
	flagsBind(ApplyCommand)

	ApplyCommand.Flags().StringP("file", "f", "", "file to apply.")
	ApplyCommand.Flags().StringP("directory", "d", "", "directory with files to apply.")

	rootCmd.AddCommand(ConvertCommand)
	rootCmd.AddCommand(VersionCommand)
	rootCmd.AddCommand(ApplyCommand)
}

func flagsBind(command *cobra.Command) {
	command.Flags().StringP("output", "o", "output", "output directory for converted templates.")
	command.Flags().StringP("templates", "t", "templates", "templates directory with files to convert with extension .kube.yaml.")
	command.Flags().StringP("env", "e", "stage", "envirement for generate deploy configuration.")
}

func Execute() {
	rootCmd.Execute()
}
