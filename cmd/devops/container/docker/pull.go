// Package docker /*
package docker

import (
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

var (
	containerName string
)

// pullCmd represents the pull command
var pullCmd = &cobra.Command{
	Use:   "pull",
	Short: "This command is used to pull the docker container from docker registry",
	Long:  `This command is used to pull the docker container from docker registry. For example:`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Please wait while the container is getting pulled ... ")
		command := exec.Command("docker", "pull", containerName)
		out, err := command.Output()
		if err != nil {
			panic(err)
		}
		fmt.Println(string(out))
	},
}

func init() {

	pullCmd.Flags().StringVarP(&containerName, "container", "c", "", "The container name to pull")
	if err := pullCmd.MarkFlagRequired("container"); err != nil {
		fmt.Println(err)
	}
	DockerCmd.AddCommand(pullCmd)
}
