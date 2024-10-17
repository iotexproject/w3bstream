package utils

import (
	"bufio"
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

const (
	smartcontractsPath = "../smartcontracts/"
)

func DeployContract(endpoint string) error {
	// A private key in Anvil local chain
	os.Setenv("PRIVATE_KEY", "0x7c852118294e51e653712a81e05800f419141751be58f605c371e15141b007a6")
	os.Setenv("URL", endpoint)

	cmd := exec.Command("bash", "-c", "./deploy.sh")
	cmd.Dir = smartcontractsPath

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return err
	}

	// Start the command asynchronously
	if err := cmd.Start(); err != nil {
		return err
	}

	scanner := bufio.NewScanner(stdout)
	scanner.Split(bufio.ScanLines)

	var outputBuffer bytes.Buffer

	for scanner.Scan() {
		line := scanner.Text()
		log.Println(line)
		outputBuffer.WriteString(line + "\n")
	}

	// Check for any scanning errors
	if err := scanner.Err(); err != nil {
		return err
	}

	if err := cmd.Wait(); err != nil {
		return err
	}

	// reprint outputbuffer
	fmt.Println("Output:")
	fmt.Println(outputBuffer.String())
	return nil
}
