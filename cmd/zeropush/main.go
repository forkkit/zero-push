package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"go.uber.org/zap/zaptest"
)

var rootCmd = &cobra.Command{
	Use:   "zeropush",
	Short: "ZeroPush is a zero-knowledge push server",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ZeroPush 😮")
		cmd.Help()
	},
}

func execute() {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_"))

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func initLogger() {
	zaptest.Level(zapcore.DebugLevel)
	config := zap.NewDevelopmentConfig()
	config.OutputPaths = []string{"stdout"}
	l, _ := config.Build()
	zap.ReplaceGlobals(l)
}

func main() {
	initLogger()
	execute()
}
