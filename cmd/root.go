package cmd

import (
	"fmt"
	"github.com/partialize/iam"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"os"
)

var (
	cfgFile string

	rootCmd = &cobra.Command{
		Use: "iam",
		RunE: func(cmd *cobra.Command, args []string) error {
			iam, err := iam.New()
			if err != nil {
				return err
			}
			return iam.Start(viper.GetString("address"))
		},
	}
)

// Execute executes the root command.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is .iam.toml)")
	rootCmd.PersistentFlags().String("address", ":8080", "server address")

	_ = viper.BindPFlag("address", rootCmd.PersistentFlags().Lookup("address"))
}

func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		viper.SetConfigFile(".iam.toml")
	}

	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
