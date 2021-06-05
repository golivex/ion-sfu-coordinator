package cmd

import (
	loadtest "github.com/manishiitg/actions/loadtest"
	"github.com/manishiitg/actions/loadtest/client/gst"
	"github.com/spf13/cobra"
)

var loadCmd = &cobra.Command{
	Use:   "loadtest",
	Short: "start the actions server",
	RunE:  loadMain,
}

var session, gaddr, file, role, loglevel, simulcast, paddr string
var total, cycle, duration int
var create_room = -1

func init() {
	loadCmd.PersistentFlags().StringVarP(&file, "file", "f", "480p", "type of file either test 360p 480p 720p h246")
	loadCmd.PersistentFlags().StringVarP(&gaddr, "gaddr", "g", "http://0.0.0.0:4000/", "SFU Cordinator")
	loadCmd.PersistentFlags().StringVarP(&session, "session", "s", "test", "join session name")
	loadCmd.PersistentFlags().IntVarP(&total, "clients", "c", 1, "Number of clients to start")
	loadCmd.PersistentFlags().IntVarP(&cycle, "cycle", "y", 1000, "Run new client cycle in ms")
	loadCmd.PersistentFlags().IntVarP(&duration, "duration", "d", 60*60, "Running duration in seconds")
	loadCmd.PersistentFlags().StringVarP(&role, "role", "r", "pubsub", "Run as pubsub/sub")
	loadCmd.PersistentFlags().IntVarP(&create_room, "create_room", "x", -1, "number of peers per room")
	rootCmd.AddCommand(loadCmd)
}

func clientThread() {

	cancel := make(chan struct{}) // not used as such
	go loadtest.Init(file, gaddr, session, total, cycle, duration, role, create_room, -1, cancel)

}

func loadMain(cmd *cobra.Command, args []string) error {
	go clientThread()
	gst.MainLoop()
	return nil
}
