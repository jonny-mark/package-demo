package main

import (
	"github.com/spf13/cobra"
	"log"
	"package-demo/cmd/model"
)

var (
	// Version is the version of the compiled software.
	Version = "v0.15.4"

	rootCmd = &cobra.Command{
		Use:     "demo",
		Short:   "demo: A microservice framework for Go",
		Long:    `demo: A microservice framework for Go`,
		Version: Version,
	}
)

func init() {
	rootCmd.AddCommand(model.CmdNew)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
