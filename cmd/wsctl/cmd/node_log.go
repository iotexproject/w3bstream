package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
)

var logCmd = &cobra.Command{
	Use:   "log",
	Short: "show node service log",
	RunE: func(cmd *cobra.Command, args []string) error {
		if _, err := os.Stat("./docker-compose.yaml"); err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrap(err, "failed to check docker-compose.yaml file stat")
			}
			return errors.Wrap(err, "docker-compose.yaml file not exist")
		}

		upCmd := exec.Command("docker-compose", "logs", "-f", "w3bnode")
		out, err := upCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to read docker-compose log")
		}
		defer out.Close()

		upCmd.Stderr = upCmd.Stdout

		if err := upCmd.Start(); err != nil {
			return errors.Wrap(err, "failed to read docker-compose log")
		}
		buff := make([]byte, 8)

		for {
			len, err := out.Read(buff)
			if err == io.EOF {
				break
			}
			fmt.Print(string(buff[:len]))
		}
		if err := upCmd.Wait(); err != nil {
			return errors.Wrap(err, "failed to read docker-compose log")
		}
		return nil
	},
}

func init() {
	nodeCmd.AddCommand(logCmd)
}
