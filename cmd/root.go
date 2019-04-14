package cmd

import (
	"github.com/spf13/cobra"

	"github.com/jace-ys/viaduct/cmd/viaduct"
	"github.com/jace-ys/viaduct/pkg/config"
	"github.com/jace-ys/viaduct/pkg/utils/log"
)

type Flags struct {
	config.Config
}

func init() {
	log.WithLevels(log.Options{
		Prefix: "Viaduct",
	})
}

func Execute() {
	cmdFlags := &Flags{}

	rootCmd := &cobra.Command{
		Use:   "viaduct",
		Short: "Viaduct is a lightweight and configurable API gateway written in Go",
	}

	startCmd := &cobra.Command{
		Use:   "start",
		Short: "Start the viaduct server",
		Run: func(cmd *cobra.Command, args []string) {
			err := setupEnv(cmdFlags)
			if err != nil {
				log.Error().Fatal(err)
			}

			err = viaduct.Start()
			if err != nil {
				log.Error().Fatal(err)
			}
		},
	}

	startCmd.Flags().StringVarP(&cmdFlags.Port, "port", "p", config.DefaultPort, "Port to run Viaduct server on")
	startCmd.Flags().StringVarP(&cmdFlags.ConfigFile, "file", "f", config.DefaultConfigFile, "Path to .yml configuration file")

	rootCmd.AddCommand(startCmd)

	err := rootCmd.Execute()
	if err != nil {
		log.Error().Fatal(err)
	}
}
