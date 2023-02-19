// Package docker /*
package docker

import (
	"github.com/spf13/cobra"
)

// Cmd dockerCmd represents the docker command
var DockerCmd = &cobra.Command{
	Use:   "docker",
	Short: "Docker is a pallet that contains docker commands",
	Long:  `You can use this pallet by (example): knife docker ps`,
	Run: func(cmd *cobra.Command, args []string) {
		cmd.Help()
	},
}

func init() {

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// dockerCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// dockerCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
