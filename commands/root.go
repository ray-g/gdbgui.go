package commands

import (
	"fmt"
	"os"
	"runtime"

	"github.com/spf13/cobra"
)

var (
	debugMode bool
)

const (
	product = "GDBGUI.Go"
	version = "0.0.1"
)

func init() {
	addGlobalFlags()
	addSubCommands()
}

var rootCmd = &cobra.Command{
	Use:   "GDBGUI.Go",
	Short: "GDBGUI.Go - GDB plus WebGUI",
	PersistentPreRun: func(cmd *cobra.Command, args []string) {
		maxProcs := runtime.GOMAXPROCS(0)
		if maxProcs < 2 {
			runtime.GOMAXPROCS(2)
		}
	},
}

func addGlobalFlags() {
	rootCmd.PersistentFlags().BoolVarP(&debugMode, "debug", "d", false, "Run with debug mode for GDBGUI.Go itself")
}

func addSubCommands() {
	rootCmd.AddCommand(
		//		serverCmd,
		//		clientCmd,
		versionCmd,
	)
}

// Execute start and go
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
