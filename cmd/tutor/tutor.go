package tutor

import (
	"fmt"
	"multiApp/internal/tutor/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var Cmd = &cobra.Command{
	Use:   "serve-http",
	Short: "Запустить сервер с приложением.",
	Long: `Сервер будет ожидать http запросы. 
`,
	RunE: run,
}

func run(cmd *cobra.Command, args []string) error {
	a, err := app.New()
	if err != nil {
		return fmt.Errorf("failed run http server: %w", err)
	}
	a.Run()

	return nil
}

func init() {
	Cmd.
		Flags().
		String("secret-key", "", "Secret key for encrypt auth token.")
	viper.BindPFlag("secret-key", Cmd.Flags().Lookup("secret-key"))
	Cmd.
		Flags().
		Uint("token-exp", 3600, "Time in seconds for jwt token would live.")
	viper.BindPFlag("token-exp", Cmd.Flags().Lookup("token-exp"))
	Cmd.
		Flags().
		String("config-file-path", "", "Path to config file")
	Cmd.MarkFlagRequired("config-file-path")
	viper.BindPFlag("config-file-path", Cmd.Flags().Lookup("config-file-path"))
}
