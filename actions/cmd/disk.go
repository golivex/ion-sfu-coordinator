package cmd

import (
	"runtime"

	"github.com/manishiitg/actions/loadtest/client/gst"
	tasktodisk "github.com/manishiitg/actions/tracktodisk"
	"github.com/spf13/cobra"
)

var diskCmd = &cobra.Command{
	Use:   "savetodisk",
	Short: "start the actions server",
	RunE:  diskMain,
}

// parse flag
var sessionf, addrf string

func init() {
	diskCmd.PersistentFlags().StringVarP(&addrf, "addr", "a", "http://0.0.0.0:4000/", "SFU Cordinator")
	diskCmd.PersistentFlags().StringVarP(&sessionf, "session", "s", "test", "session to join")
	rootCmd.AddCommand(diskCmd)
}

func compositeThread() {
	cancel := make(chan struct{})
	tasktodisk.Init(sessionf, addrf, "webm", cancel)
}

func diskMain(cmd *cobra.Command, args []string) error {
	runtime.LockOSThread()
	go compositeThread()
	gst.MainLoop()
	return nil
}
