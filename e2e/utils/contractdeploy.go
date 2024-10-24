package utils

import (
	"bufio"
	"bytes"
	"log"
	"os"
	"os/exec"
	"regexp"
)

const (
	smartcontractsPath = "../smartcontracts/"
)

var (
	taskManagerRe      = regexp.MustCompile(`W3bstreamTaskManager deployed to (\S+)`)
	proverRe           = regexp.MustCompile(`W3bstreamProver deployed to (\S+)`)
	daoRe              = regexp.MustCompile(`W3bstreamDAO deployed to (\S+)`)
	minterRe           = regexp.MustCompile(`W3bstreamBlockMinter deployed to (\S+)`)
	projectRegistrarRe = regexp.MustCompile(`ProjectRegistrar deployed to (\S+)`)
	mockProjectRe      = regexp.MustCompile(`MockProject deployed to (\S+)`)
	wsProjectRe        = regexp.MustCompile(`W3bstreamProject deployed to (\S+)`)
	routerRe           = regexp.MustCompile(`W3bstreamRouter deployed to (\S+)`)
	mockDappRe         = regexp.MustCompile(`MockProcessor deployed to (\S+)`)
	fleetManagementRe  = regexp.MustCompile(`FleetManagement deployed to (\S+)`)
)

type ContractsDeployments struct {
	TaskManager     string
	Prover          string
	Dao             string
	Minter          string
	Registrar       string
	MockProject     string
	WSProject       string
	Router          string
	MockDapp        string
	FleetManagement string
}

func DeployContract(endpoint string, payerHex string) (*ContractsDeployments, error) {
	os.Setenv("PRIVATE_KEY", payerHex)
	os.Setenv("URL", endpoint)

	cmd := exec.Command("bash", "-c", "./deploy.sh --network dev")
	cmd.Dir = smartcontractsPath

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	// Start the command asynchronously
	if err := cmd.Start(); err != nil {
		return nil, err
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
		return nil, err
	}

	if err := cmd.Wait(); err != nil {
		return nil, err
	}

	output := outputBuffer.String()
	deployments := &ContractsDeployments{}

	// Match each line against the regex patterns
	if match := taskManagerRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.TaskManager = match[1]
	}
	if match := proverRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.Prover = match[1]
	}
	if match := daoRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.Dao = match[1]
	}
	if match := minterRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.Minter = match[1]
	}
	if match := projectRegistrarRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.Registrar = match[1]
	}
	if match := mockProjectRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.MockProject = match[1]
	}
	if match := wsProjectRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.WSProject = match[1]
	}
	if match := mockDappRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.MockDapp = match[1]
	}
	if match := routerRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.Router = match[1]
	}
	if match := fleetManagementRe.FindStringSubmatch(output); len(match) > 1 {
		deployments.FleetManagement = match[1]
	}

	return deployments, nil
}
