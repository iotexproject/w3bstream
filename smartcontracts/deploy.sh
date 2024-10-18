#!/usr/bin/env bash

# Ensure the script stops on errors
set -e

# Optional: Navigate to the project directory (if needed)
# cd /path/to/your/project

# Check if Yarn is installed
if ! command -v yarn &> /dev/null; then
    echo "Yarn is not installed. Please install it first."
    exit 1
fi

# Check if Hardhat is installed in the project
if ! yarn list --pattern hardhat | grep -q 'hardhat'; then
    echo "Hardhat is not installed. Installing Hardhat..."
    yarn add --dev hardhat
fi

# Check if the PRIVATE_KEY environment variable is set
if [[ -z "${PRIVATE_KEY}" ]]; then
    echo "Error: PRIVATE_KEY environment variable is not set."
    exit 1
fi

# Run the Hardhat deployment script
echo "Running Hardhat deployment..."
yarn hardhat run scripts/deploy.ts --network dev

# Check if the previous command was successful
if [ $? -eq 0 ]; then
    echo "Deployment completed successfully."
else
    echo "Deployment failed."
    exit 1
fi
