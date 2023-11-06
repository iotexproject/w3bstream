package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

// downCmd represents the down command
var downCmd = &cobra.Command{
	Use:   "down",
	Short: "shutdown the node service",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("./docker-compose.yaml"); err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrap(err, "failed to check docker-compose.yaml file stat")
			}
			return errors.Wrap(err, "docker-compose.yaml file not exist")
		}

		downCmd := exec.Command("docker-compose", "down")
		out, err := downCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to stop docker-compose")
		}
		defer out.Close()

		downCmd.Stderr = downCmd.Stdout

		if err := downCmd.Start(); err != nil {
			return errors.Wrap(err, "failed to stop docker-compose")
		}
		buff := make([]byte, 8)

		for {
			len, err := out.Read(buff)
			if err == io.EOF {
				break
			}
			fmt.Print(string(buff[:len]))
		}
		if err := downCmd.Wait(); err != nil {
			return errors.Wrap(err, "failed to stop docker-compose")
		}
		return nil
	},
}

func init() {
	nodeCmd.AddCommand(downCmd)
}
