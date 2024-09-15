package tutor

import (
	"multiApp/internal/tutor/app"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var TutorCmd = &cobra.Command{
	Use:   "serve-http",
	Short: "Запустить сервер с приложением.",
	Long: `Сервер будет ожидать http запросы. 
`,
	Run: run,
}

func run(cmd *cobra.Command, args []string) {
	a := app.New()
	a.Run()
}

func init() {
	TutorCmd.
		Flags().
		String("secret-key", "", "Secret key for encrypt auth token.")
	TutorCmd.MarkFlagRequired("secret-key")
	viper.BindPFlag("secret-key", TutorCmd.Flags().Lookup("secret-key"))
}
