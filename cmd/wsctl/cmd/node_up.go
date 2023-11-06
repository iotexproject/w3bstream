package cmd

import (
	"fmt"
	"io"
	"os"
	"os/exec"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var upCmd = &cobra.Command{
	Use:   "up",
	Short: "startup the node service",
	RunE: func(cmd *cobra.Command, args []string) error {
		privateKey, err := cmd.Flags().GetString("private-key")
		if err != nil {
			return errors.Wrap(err, "failed to get flag private-key")
		}
		projectFile, err := cmd.Flags().GetString("project-file")
		if err != nil {
			return errors.Wrap(err, "failed to get flag project-file")
		}
		if projectFile == "" {
			projectFile = "test/data/create.json"
		}

		_, err = os.Stat("./docker-compose.yaml")
		if err != nil {
			if !os.IsNotExist(err) {
				return errors.Wrap(err, "failed to check docker-compose.yaml file stat")
			}

			// TODO auto download docker-compose.yaml when not exist
			return errors.Wrap(err, "docker-compose.yaml file not exist")
		}
		if err := createEnvFile(viper.Get(NodePort).(string), projectFile, privateKey); err != nil {
			return err
		}

		upCmd := exec.Command("docker-compose", "up", "-d")
		out, err := upCmd.StdoutPipe()
		if err != nil {
			return errors.Wrap(err, "failed to start docker-compose")
		}
		defer out.Close()

		upCmd.Stderr = upCmd.Stdout

		if err := upCmd.Start(); err != nil {
			return errors.Wrap(err, "failed to start docker-compose")
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
			return errors.Wrap(err, "failed to start docker-compose")
		}
		return nil
	},
}

func createEnvFile(port, projectFile, privateKey string) error {
	out, err := os.Create("./.env")
	if err != nil {
		return errors.Wrap(err, "failed to create env file")
	}
	defer out.Close()

	fmt.Fprintf(out, "%s=%s\n", "PORT", port)
	fmt.Fprintf(out, "%s=%s\n", "PROJECT_CONFIG_FILE", projectFile)
	fmt.Fprintf(out, "%s=%s\n", "OPERATOR_PRIVATE_KEY", privateKey)
	return nil
}

func init() {
	nodeCmd.AddCommand(upCmd)

	upCmd.Flags().StringP("private-key", "p", "", "blockchain operate private key")
	upCmd.Flags().StringP("project-file", "f", "", "project file full path")

	upCmd.MarkFlagRequired("private-key")
}
