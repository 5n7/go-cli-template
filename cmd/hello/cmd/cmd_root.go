package cmd

import (
	"github.com/skmatz/go-cli-template/cli"
	"github.com/spf13/cobra"
)

var (
	opt     cli.Options
	version bool
)

func runRoot(cmd *cobra.Command, args []string) error {
	if version {
		return runVersion(cmd, args)
	}

	c := cli.New()
	return c.Run(opt)
}

var rootCmd = &cobra.Command{
	Use:   "hello",
	Short: "Say hello",
	Long:  "Say hello.",
	RunE:  runRoot,
}

func init() {
	rootCmd.Flags().StringVarP(&opt.Name, "name", "n", "", "name")
	rootCmd.Flags().BoolVarP(&version, "version", "V", false, "show version")
}

func Execute() {
	rootCmd.Execute() //nolint:errcheck
}
