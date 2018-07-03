package commands

import (
	"fmt"

	"github.com/spf13/cobra"
)

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version of GDBGUI.Go",
	Long:  `All software has versions. This is GDBGUI.Go's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("GDBGUI.Go - v%s\n", version)
	},
}
