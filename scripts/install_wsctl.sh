#!/bin/bash

# GitHub API URL for the latest release of machinefi/sprout
GITHUB_API_URL="https://api.github.com/repos/machinefi/sprout/releases/latest"

# Detect the operating system
OS="$(uname)"

# Function to download wsctl based on OS
download_wsctl() {
    local os_name=$1
    local binary_name="wsctl-${os_name}-amd64"
    
    if [ "$os_name" = "windows" ]; then
        binary_name="${binary_name}.exe"
    fi

    local download_url=$(curl -s $GITHUB_API_URL | grep "browser_download_url.*$binary_name" | cut -d '"' -f 4 | head -n 1)

    if [ -z "$download_url" ]; then
        echo "Download URL not found for $binary_name"
        exit 1
    fi

    echo "Downloading wsctl for ${os_name}..."
    curl -L $download_url -o wsctl
    chmod +x wsctl

    echo "wsctl downloaded and made executable."
}

case "$OS" in
    'Linux')
        download_wsctl "linux"
        ;;
    'Darwin')
        download_wsctl "darwin"
        ;;
    'MINGW'*|'MSYS'*|'CYGWIN'*)
        download_wsctl "windows"
        ;;
    *)
        echo "Unsupported OS: $OS"
        exit 1
        ;;
esac

# Move wsctl to a directory in PATH (optional, modify as needed)
# Note: This step may not be applicable or may differ for Windows
sudo mv wsctl /usr/local/bin/

# Any additional setup or configuration steps


# Output the version of wsctl
echo "wsctl installed successfully."
/usr/local/bin/wsctl version
