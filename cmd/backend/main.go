package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"

	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/handler"
	"github.com/Pineapple217/cvrs/pkg/server"
	"github.com/spf13/cobra"
)

const version = "v0.0.0"
const bannerTemplate = `
 ██████ ██    ██ ██████  ███████ 
██      ██    ██ ██   ██ ██      
██      ██    ██ ██████  ███████ 
██       ██  ██  ██   ██      ██ 
 ██████   ████   ██   ██ ███████ %s

https://github.com/Pineapple217/cvrs
-----------------------------------------------------------------------------`

func main() {
	slog.SetDefault(slog.New(slog.Default().Handler()))
	banner := fmt.Sprintf(bannerTemplate, version)

	cmdRun := &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(banner)
			os.Stdout.Sync()

			db := database.Database{}

			h := handler.NewHandler(&db)

			server := server.NewServer()
			server.RegisterRoutes(h)
			server.ApplyMiddleware()
			server.Start()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit
			slog.Info("Received an interrupt signal, exiting...")

			server.Stop()
		},
	}

	var rootCmd = &cobra.Command{Use: "cvrs"}
	rootCmd.Version = version
	rootCmd.AddCommand(cmdRun)
	rootCmd.Execute()
}
