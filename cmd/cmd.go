package cmd

import (
	"errors"
	"fmt"
	"ginstart/cmd/start"
	"ginstart/cmd/version"
	"os"

	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:          `ginstart`,
		Short:        `ginstart`,
		Long:         `ginstart`,
		SilenceUsage: true,
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				tip()
				return errors.New("requires at least one arg")
			}
			return nil
		},
		PersistentPreRunE: func(*cobra.Command, []string) error { return nil },
		Run: func(cmd *cobra.Command, args []string) {
			tip()
		},
	}
)

func tip() {
	usageStr := `欢迎使用gincli 帮助命令:-h`
	fmt.Printf("%s\n", usageStr)
}

func init() {
	rootCmd.AddCommand(start.StartCmd)
	rootCmd.AddCommand(version.StartCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(-1)
	}
}
