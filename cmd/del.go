package cmd

import (
	"errors"
	"fmt"
	"github.com/ythereewater/agenda-go/service"
	"github.com/ythereewater/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/ythereewater/agenda-go/logs"
)

var delCmd = &cobra.Command{
	Use:   "del",
	Short: "delete current account",
	Long:  `delete current account`,
	Run: func(cmd *cobra.Command, args []string) {
		password, _ := cmd.Flags().GetString("password")
		if password == "" {
			tools.Report(errors.New("password required"))
		}
		if err := service.DeleteUser(password); err == nil {
			fmt.Println("Success")
			logs.EventLog("delete a user")
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(delCmd)

	delCmd.Flags().StringP("password", "p", "", "your password")
}
