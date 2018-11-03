package cmd

import (
	"errors"
	"fmt"
	"github.com/ythereewater/agenda-go/service"
	"github.com/ythereewater/agenda-go/tools"
	"github.com/spf13/cobra"
	"github.com/ythereewater/agenda-go/logs"
)


var quitCmd = &cobra.Command{
	Use:   "quit",
	Short: "quit a meeting",
	Long:  `quit a meeting`,
	Run: func(cmd *cobra.Command, args []string) {
		title, _ := cmd.Flags().GetString("title")
		if title == "" {
			tools.Report(errors.New("title required"))
		}
		err := service.QuitMeeting(title)
		if err == nil {
			fmt.Println("Success")
			logs.EventLog("quit meeting: " + title)
		} else {
			tools.Report(err)
		}
	},
}

func init() {
	RootCmd.AddCommand(quitCmd)

	quitCmd.Flags().StringP("title", "t", "", "title of the meeting")
}
