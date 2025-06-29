package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"runtime/pprof"

	"github.com/Pineapple217/cvrs/pkg/config"
	"github.com/Pineapple217/cvrs/pkg/database"
	"github.com/Pineapple217/cvrs/pkg/handler"
	"github.com/Pineapple217/cvrs/pkg/server"
	"github.com/Pineapple217/cvrs/pkg/users"
	"github.com/Pineapple217/cvrs/pkg/util"
	"github.com/Pineapple217/cvrs/pkg/worker"
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

var (
	enableProfile bool
	profFile      *os.File
)

func main() {
	slog.SetDefault(slog.New(slog.Default().Handler()))
	banner := fmt.Sprintf(bannerTemplate, version)

	cmdRun := &cobra.Command{
		Use: "run",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println(banner)
			os.Stdout.Sync()

			conf, err := config.Load()
			util.MaybeDie(err, "Failed to laod config")

			db, err := database.NewDatabase(conf.Database)
			util.MaybeDieErr(err)

			wf := worker.NewWorkforce(conf.Workforce, db)
			err = wf.Start()
			util.MaybeDie(err, "Failed to start workforce")
			defer wf.Stop()

			h := handler.NewHandler(db)

			server := server.NewServer()
			server.RegisterRoutes(h)
			server.ApplyMiddleware(true)
			server.Start()
			defer server.Stop()

			quit := make(chan os.Signal, 1)
			signal.Notify(quit, os.Interrupt)
			<-quit
			slog.Info("Received an interrupt signal, exiting...")
		},
	}
	var rootCmd = &cobra.Command{
		Use:     "cvrs",
		Version: version,
		PersistentPreRunE: func(cmd *cobra.Command, args []string) error {
			if enableProfile {
				slog.Info("Running with cpu profiler")
				var err error
				profFile, err = os.Create("cpu.prof")
				if err != nil {
					return fmt.Errorf("could not create CPU profile: %w", err)
				}
				pprof.StartCPUProfile(profFile)
			}
			return nil
		},
		PersistentPostRun: func(cmd *cobra.Command, args []string) {
			if enableProfile && profFile != nil {
				pprof.StopCPUProfile()
				profFile.Close()
				slog.Debug("CPU profile saved", "file", "cpu.prof")
			}
		},
	}
	rootCmd.PersistentFlags().BoolVar(&enableProfile, "profile", false, "Enable CPU profiling and write to cpu.prof")

	rootCmd.AddCommand(cmdRun)
	rootCmd.AddCommand(users.GetCmd())
	if err := rootCmd.Execute(); err != nil {
		slog.Error("Command execution failed", "error", err)
		os.Exit(1)
	}
}
