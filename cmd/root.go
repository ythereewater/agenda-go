package cmd

import (
	"fmt"
	"os"
	"github.com/spf13/cobra"
	"github.com/ythereewater/agenda-go/logs"
)

var cfgFile string


var RootCmd = &cobra.Command{
	Use:   "agd",
	Short: "Agenda-go",
	Long:  "A tool for managing meetings",
}

func Execute() {
	if err := RootCmd.Execute(); err != nil {
		fmt.Println(err)
		logs.ErrLog(err)
		os.Exit(1)
	}
}
