/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>

*/
package docker

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var container string

// pushCmd represents the push command
var pushCmd = &cobra.Command{
	Use:   "push",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please wait while the container is getting pushed ... ")
		command := exec.Command("docker", "push", container)
		out, err := command.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	},
}

func init() {
	pushCmd.Flags().StringVarP(&container, "container", "c", "", "The container name to push")
	if err := pushCmd.MarkFlagRequired("container"); err != nil {
		fmt.Println(err)
	}
	DockerCmd.AddCommand(pushCmd)
}
