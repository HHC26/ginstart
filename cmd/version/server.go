package version

import (
	"fmt"

	"github.com/spf13/cobra"
)

var Version = "V1.0.0"

var (
	StartCmd = &cobra.Command{
		Use:     "version",
		Short:   "Get version info",
		Example: "ginstart version",
		PreRun: func(cmd *cobra.Command, args []string) {

		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return run()
		},
	}
)

func run() error {
	fmt.Println(Version)
	return nil
}
