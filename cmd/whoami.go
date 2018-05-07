package cmd

import (
	"fmt"
	"os/user"

	"github.com/spf13/cobra"
)

var (
	whoamiCmd = &cobra.Command{
		Use:   "whoami",
		Short: "return current user",
		Long:  "",
		Run: func(cmd *cobra.Command, args []string) {
			Whoami()
		},
	}
)

func init() {
	RootCmd.AddCommand(whoamiCmd)
}

func Whoami() {
	usr, err := user.Current()
	if err != nil {
		fmt.Printf("cugo: %s", err)
		return
	}

	fmt.Printf("%s\n", usr.Username)
	return
}
