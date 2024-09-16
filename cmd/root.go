/*
Copyright © 2024 Ilya Otinov <ilya.otinov@gmail.com>
*/
package cmd

import (
	"multiApp/cmd/tutor"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "subt",
	Short: "Платформа, где люди могут обучать друг друга.",
	Long: ` Здесь учителя и ученики могут эффективно взаимодействовать.
А так же учителя делиться материалами друг с другом.
`,
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	rootCmd.AddCommand(
		tutor.Cmd,
	)
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
}
