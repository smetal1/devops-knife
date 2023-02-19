/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package docker

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// psCmd represents the ps command
var psCmd = &cobra.Command{
	Use:   "ps",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please wait while all running container in the system is getting fetched ... ")
		command := exec.Command("docker", "ps")
		out, err := command.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	DockerCmd.AddCommand(psCmd)
}
