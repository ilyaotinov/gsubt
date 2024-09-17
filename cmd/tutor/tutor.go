package tutor

import (
	"fmt"
	"multiApp/internal/tutor/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

const (
	defaultJwtTokenExpInSeconds = 3600
	configFilePathOption        = "config-file-path"
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
	_ = viper.BindPFlag("secret-key", Cmd.Flags().Lookup("secret-key"))
	Cmd.
		Flags().
		Uint("token-exp", defaultJwtTokenExpInSeconds, "Time in seconds for jwt token would live.")
	_ = viper.BindPFlag("token-exp", Cmd.Flags().Lookup("token-exp"))
	Cmd.
		Flags().
		String(configFilePathOption, "", "Path to config file")
	_ = Cmd.MarkFlagRequired(configFilePathOption)
	_ = viper.BindPFlag(configFilePathOption, Cmd.Flags().Lookup(configFilePathOption))
}
