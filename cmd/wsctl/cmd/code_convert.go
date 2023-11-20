package cmd

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"io/fs"
	"os"
)

var convertCmd = &cobra.Command{
	Use:   "convert",
	Short: "convert zkp code to hex string compressed with zlib",
	RunE: func(cmd *cobra.Command, args []string) error {
		vmType, err := cmd.Flags().GetString("vm-type")
		if err != nil {
			return errors.Wrap(err, "failed to get flag vm-type")
		}
		codeFile, err := cmd.Flags().GetString("code-file")
		if err != nil {
			return errors.Wrap(err, "failed to get flag code-file")
		}
		confFile, err := cmd.Flags().GetString("conf-file")
		if err != nil {
			return errors.Wrap(err, "failed to get flag conf-file")
		}
		expParam, err := cmd.Flags().GetString("expand-param")
		if err != nil {
			return errors.Wrap(err, "failed to get flag expand-param")
		}

		content, err := os.ReadFile(codeFile)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to read code-file %s", codeFile))
		}

		var b bytes.Buffer
		w := zlib.NewWriter(&b)
		defer w.Close()
		_, err = w.Write(content)
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to zlib code-file %s", codeFile))
		}

		hexString := hex.EncodeToString(b.Bytes())

		confMap := make(map[string]interface{})
		if expParam != "" {
			confMap["codeExpParam"] = expParam
		}
		confMap["vmType"] = vmType
		confMap["code"] = hexString

		jsonConf, err := json.MarshalIndent(confMap, "", "  ")

		if confFile == "" {
			confFile = fmt.Sprintf("./%s-config.json", vmType)
		}
		err = os.WriteFile(confFile, jsonConf, fs.FileMode(0666))
		if err != nil {
			return errors.Wrap(err, fmt.Sprintf("failed to write conf-file %s", confFile))
		}

		return nil
	},
}

func init() {
	codeCmd.AddCommand(convertCmd)

	convertCmd.Flags().StringP("vm-type", "t", "", "vm type")
	convertCmd.Flags().StringP("code-file", "i", "", "the zkp code file")
	convertCmd.Flags().StringP("conf-file", "o", "", "the config file")
	convertCmd.Flags().StringP("expand-param", "e", "", "the parameters that need to be expanded")

	convertCmd.MarkFlagRequired("vm-type")
	convertCmd.MarkFlagRequired("code-file")
	//convertCmd.MarkFlagRequired("conf-file")
	//convertCmd.MarkFlagRequired("expand-param")
}
