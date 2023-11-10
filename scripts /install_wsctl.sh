#!/bin/bash

# Set the GitHub API URL for the releases of wsctl
GITHUB_API_URL="https://api.github.com/repos/machinefi/w3bstream-mainnet/releases/latest"

# Use curl to get the download URL for the latest release of wsctl
DOWNLOAD_URL=$(curl -s $GITHUB_API_URL | grep "browser_download_url" | cut -d '"' -f 4 | head -n 1)

# Check if the download URL is empty
if [ -z "$DOWNLOAD_URL" ]; then
    echo "Error: Could not find the download URL for wsctl."
    exit 1
fi

# Download the latest release of wsctl
curl -L $DOWNLOAD_URL -o wsctl

# Make the wsctl file executable
chmod +x wsctl

# Move wsctl to /usr/local/bin so it can be run from anywhere
sudo mv wsctl /usr/local/bin/

# Output the version of wsctl
echo "wsctl installed successfully."
/usr/local/bin/wsctl --version
