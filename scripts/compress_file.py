import os
import zlib
import binascii
import ipfshttpclient
import sys

dns_endpoint = 'ipfs.mainnet.iotex.io'
def convert_code_to_zlib_hex(code_file: str) -> str:
    try:
        with open(code_file, 'rb') as f:
            content = f.read()

        compressed_data = zlib.compress(content)

        hex_string = binascii.hexlify(compressed_data).decode('utf-8')

        return hex_string
    except Exception as e:
        raise RuntimeError(f"Failed to convert and compress the code file {code_file}") from e

def upload_to_ipfs(hex_string: str, endpoint: str = f"/dns/{dns_endpoint}/tcp/443/https") -> str:
    try:
        client = ipfshttpclient.connect(endpoint)

        res = client.add_bytes(hex_string.encode('utf-8'))
        cid = res

        client.pin.add(cid)

        return f"ipfs://{dns_endpoint}/{cid}"
    except Exception as e:
        raise RuntimeError(f"Failed to upload file to IPFS: {e}")

if __name__ == '__main__':
    if len(sys.argv) != 2:
        print("Usage: python3 compress_file.py <code_file>")
        sys.exit(1)

    code_file = sys.argv[1]

    # Convert code file to zlib compressed hex string
    hex_string = convert_code_to_zlib_hex(code_file)
    print(f"Zlib Hex String: {hex_string}")

    # Upload to IPFS
    ipfs_hash = upload_to_ipfs(hex_string)
    print(f"IPFS Hash: {ipfs_hash}")